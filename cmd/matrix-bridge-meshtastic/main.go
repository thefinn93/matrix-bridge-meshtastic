package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"

	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/config"
	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/meshtastic"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	logrus.SetLevel(logrus.DebugLevel)
	if err := config.Load(); err != nil {
		return err
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go meshtastic.Receive(ctx)

	if err := meshtastic.SendConfigInit(ctx); err != nil {
		logrus.WithError(err).Error("error sending init to meshtastic")
	}

	<-ctx.Done()

	meshtastic.ShutdownReceiver()

	return nil
}
