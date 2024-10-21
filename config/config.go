package config

import (
	"encoding/json"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"maunium.net/go/mautrix/id"
)

type Config struct {
	Meshtastic Meshtastic
	Matrix     Matrix
	Database   string
}

type Meshtastic struct {
	Address         string
	RequestTimeout  time.Duration
	PollingInterval time.Duration
}

type Matrix struct {
	User     id.UserID
	Password string
	Room     id.RoomID
}

var C = Config{
	Database: "matrix-bridge-meshtastic.db",
	Meshtastic: Meshtastic{
		RequestTimeout:  time.Second * 5,
		PollingInterval: time.Millisecond * 500,
	},
}

func Load() error {
	for _, filename := range []string{"/etc/matrix-bridge-meshtastic.json", "matrix-bridge-meshtastic.json"} {
		err := load(filename)
		if err != nil && !os.IsNotExist(err) {
			return err
		}
	}

	return nil
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
