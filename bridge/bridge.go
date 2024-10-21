package bridge

import (
	"context"
	"fmt"

	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/matrix"
	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/meshtastic/protobufs"
	"github.com/sirupsen/logrus"
)

func RunBridge(ctx context.Context, fromRadioCh chan *protobufs.FromRadio) {
	for {
		fromRadio := <-fromRadioCh
		switch payload := fromRadio.PayloadVariant.(type) {
		case *protobufs.FromRadio_Channel:
			logrus.WithField("type", "channel").Debugf("received %+v", payload)
		case *protobufs.FromRadio_Packet:
			if err := handlePacket(ctx, payload.Packet); err != nil {
				logrus.WithError(err).Error("error handling meshtastic packet")
			}
		default:
			logrus.Debugf("received unknown message type: %+v", payload)
		}
	}
}

func handlePacket(ctx context.Context, packet *protobufs.MeshPacket) error {
	payload, ok := packet.PayloadVariant.(*protobufs.MeshPacket_Decoded)
	if !ok {
		return nil // ignore encrypted packets for now
	}

	switch payload.Decoded.Portnum {
	case protobufs.PortNum_TEXT_MESSAGE_APP:
		matrix.SendMessage(ctx, fmt.Sprintf("text from %x: %s (snr: %f, rssi: %d, hop limit: %d, hop start: %d)", packet.From, payload.Decoded.Payload, packet.RxSnr, packet.RxRssi, packet.HopLimit, packet.HopStart))
	default:
		logrus.WithField("type", protobufs.PortNum_name[int32(payload.Decoded.Portnum)]).Debug("ignoring unknown app payload")
	}

	return nil
}
