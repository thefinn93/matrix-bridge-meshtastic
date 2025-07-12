package bridge

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"git.janky.solutions/finn/matrix-bridge-meshtastic/db"
	"git.janky.solutions/finn/matrix-bridge-meshtastic/matrix"
	"git.janky.solutions/finn/matrix-bridge-meshtastic/meshtastic/protobufs"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

var channelNames = map[int32]string{}

func RunBridge(ctx context.Context, fromRadioCh chan *protobufs.FromRadio) {
	for {
		var err error
		fromRadio := <-fromRadioCh
		switch payload := fromRadio.PayloadVariant.(type) {
		case *protobufs.FromRadio_Channel:
			logrus.WithField("type", "channel").Debugf("received %+v", payload)
			channelNames[payload.Channel.Index] = payload.Channel.Settings.Name
		case *protobufs.FromRadio_NodeInfo:
			logrus.WithField("type", "node_info").Debugf("received %+v", payload)
			err = handleNodeInfo(ctx, payload.NodeInfo)
		case *protobufs.FromRadio_Packet:
			err = handlePacket(ctx, payload.Packet)
		default:
			logrus.Debugf("received unknown message type: %+v", payload)
		}

		if err != nil {
			logrus.WithError(err).Error("error handling meshtastic message")
		}
	}
}

func handleNodeInfo(ctx context.Context, nodeInfo *protobufs.NodeInfo) error {
	if nodeInfo == nil {
		logrus.Warn("handleNodeInfo called with null nodeInfo")
		return nil
	}

	queries, dbconn, err := db.Get()
	if err != nil {
		return err
	}
	defer dbconn.Close()

	params := db.MeshtasticNodeUpdateParams{NodeNum: int64(nodeInfo.Num)}

	if nodeInfo.User != nil { // sometimes this is null, it seems
		params.MeshtasticID = nodeInfo.User.Id
		params.LongName = db.NullableString(nodeInfo.User.LongName)
		params.ShortName = db.NullableString(nodeInfo.User.ShortName)
		params.HwModel = db.NullableString(nodeInfo.User.HwModel.String())
		params.PublicKey = nodeInfo.User.PublicKey
	}

	err = queries.MeshtasticNodeUpdate(ctx, params)
	if err != nil {
		return err
	}

	logrus.WithField("node_num", nodeInfo.Num).Debug("updated info for node")

	return nil
}

func handlePacket(ctx context.Context, packet *protobufs.MeshPacket) error {
	payload, ok := packet.PayloadVariant.(*protobufs.MeshPacket_Decoded)
	if !ok {
		return nil // ignore encrypted packets for now
	}

	queries, dbconn, err := db.Get()
	if err != nil {
		return err
	}
	defer dbconn.Close()

	sourceNode, err := queries.MeshtasticNodeGet(ctx, int64(packet.From))
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logrus.WithError(err).Error("error getting packet source node from database")
	}

	var sourceString string
	if sourceNode.LongName.Valid {
		sourceString = fmt.Sprintf("%s (%s/%x)", sourceNode.LongName.String, sourceNode.ShortName.String, packet.From)
	} else {
		sourceString = fmt.Sprintf("%x", packet.From)
	}

	log := logrus.WithFields(logrus.Fields{
		"app":        protobufs.PortNum_name[int32(payload.Decoded.Portnum)],
		"source":     sourceString,
		"dest":       payload.Decoded.Dest,
		"request_id": payload.Decoded.RequestId,
		"reply_id":   payload.Decoded.ReplyId,
	})

	switch payload.Decoded.Portnum {
	case protobufs.PortNum_TEXT_MESSAGE_APP:
		log.Info("handling incoming meshtastic -> matrix message")
		channelName := channelNames[int32(packet.Channel)]
		msg := fmt.Sprintf("%s: %s (snr: %f, rssi: %d, hop: %d/%d, channel: %s)", sourceString, payload.Decoded.Payload, packet.RxSnr, packet.RxRssi, packet.HopStart-packet.HopLimit, packet.HopStart, channelName)
		return matrix.SendMessage(ctx, msg)

	case protobufs.PortNum_NODEINFO_APP:
		user := &protobufs.User{}
		err = proto.Unmarshal(payload.Decoded.Payload, user)
		if err != nil {
			return fmt.Errorf("error parsing nodeinfo app packet: %v", err)
		}

		log.Debugf("%+v", user)

	case protobufs.PortNum_TELEMETRY_APP:
		telemetry := &protobufs.Telemetry{}
		err = proto.Unmarshal(payload.Decoded.Payload, telemetry)
		if err != nil {
			return fmt.Errorf("error parsing telemetry app packet: %v", err)
		}

		log.Debugf("%+v", telemetry)

	case protobufs.PortNum_POSITION_APP:
		position := &protobufs.Position{}
		err = proto.Unmarshal(payload.Decoded.Payload, position)
		if err != nil {
			return fmt.Errorf("error parsing telemetry app packet: %v", err)
		}

		log.Debugf("%+v", position)

	default:
		log.WithFields(logrus.Fields{
			"emoji":         payload.Decoded.Emoji,
			"payload":       string(payload.Decoded.Payload),
			"payload_bytes": fmt.Sprintf("%x", payload.Decoded.Payload),
		}).Debug("ignoring unknown app")
	}

	return nil
}
