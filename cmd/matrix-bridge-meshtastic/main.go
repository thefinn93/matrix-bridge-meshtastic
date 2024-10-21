package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"

	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/bridge"
	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/config"
	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/db"
	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/matrix"
	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/meshtastic"
	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/meshtastic/protobufs"
)

func main() {
	if err := run(); err != nil {
		logrus.WithError(err).Fatal("error running bridge")
	}
}

func run() error {
	logrus.SetLevel(logrus.DebugLevel)
	if err := config.Load(); err != nil {
		return err
	}

	if err := db.Migrate(); err != nil {
		return err
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := matrix.Setup(ctx); err != nil {
		return err
	}

	fromRadioCh := make(chan *protobufs.FromRadio)
	go matrix.Run(ctx)
	go meshtastic.Receive(ctx, fromRadioCh)
	go bridge.RunBridge(ctx, fromRadioCh)

	if err := meshtastic.SendConfigInit(ctx); err != nil {
		logrus.WithError(err).Error("error sending init to meshtastic")
	}

	<-ctx.Done()

	meshtastic.ShutdownReceiver()
	matrix.Shutdown()

	return nil
}
