package config

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Meshtastic Meshtastic
}

type Meshtastic struct {
	Address string
}

var C Config

func Load() error {
	return load("matrix-meshtastic-bridge.json")
}

func load(filename string) error {
	logrus.WithField("file", filename).Info("reading config file")

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&C)
	if err != nil {
		return err
	}

	return nil
}
