package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"

	"git.janky.solutions/finn/matrix-bridge-meshtastic/bridge"
	"git.janky.solutions/finn/matrix-bridge-meshtastic/config"
	"git.janky.solutions/finn/matrix-bridge-meshtastic/db"
	"git.janky.solutions/finn/matrix-bridge-meshtastic/matrix"
	"git.janky.solutions/finn/matrix-bridge-meshtastic/meshtastic"
	"git.janky.solutions/finn/matrix-bridge-meshtastic/meshtastic/protobufs"
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

	logrus.Info("migrating database")
	if err := db.Migrate(); err != nil {
		return err
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := matrix.Setup(ctx); err != nil {
		return err
	}

	fromRadioCh := make(chan *protobufs.FromRadio, 100)
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
