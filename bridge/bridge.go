package bridge

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/db"
	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/matrix"
	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/meshtastic/protobufs"
	"github.com/sirupsen/logrus"
)

func RunBridge(ctx context.Context, fromRadioCh chan *protobufs.FromRadio) {
	for {
		var err error
		fromRadio := <-fromRadioCh
		switch payload := fromRadio.PayloadVariant.(type) {
		case *protobufs.FromRadio_Channel:
			logrus.WithField("type", "channel").Debugf("received %+v", payload)
		case *protobufs.FromRadio_NodeInfo:
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

	logrus.WithField("source", sourceString).Debug("handling incoming meshtastic node")

	switch payload.Decoded.Portnum {
	case protobufs.PortNum_TEXT_MESSAGE_APP:
		return matrix.SendMessage(ctx, fmt.Sprintf("text from %s: %s (snr: %f, rssi: %d, hop limit: %d, hop start: %d)", sourceString, payload.Decoded.Payload, packet.RxSnr, packet.RxRssi, packet.HopLimit, packet.HopStart))
	default:
		logrus.WithField("type", protobufs.PortNum_name[int32(payload.Decoded.Portnum)]).Debug("ignoring unknown app payload")
	}

	return nil
}
