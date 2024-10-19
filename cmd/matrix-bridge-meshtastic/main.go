package main

import (
	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/config"
	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/meshtastic"
	"github.com/sirupsen/logrus"
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

	return meshtastic.Receive()
}
