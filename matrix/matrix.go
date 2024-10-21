package matrix

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/config"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"maunium.net/go/mautrix"
)

var client *mautrix.Client

func Setup(ctx context.Context) error {
	clientWellKnown, err := mautrix.DiscoverClientAPI(ctx, config.C.Matrix.User.Homeserver())
	if err != nil {
		return err
	}

	homeserverURL, err := url.Parse(clientWellKnown.Homeserver.BaseURL)
	if err != nil {
		return fmt.Errorf("error parsing homeserver URL %s: %v", clientWellKnown.Homeserver.BaseURL, err)
	}

	client = &mautrix.Client{
		HomeserverURL: homeserverURL,
		UserAgent:     "matrix-meshtastic-bridge-go/unversioned",
		Client:        &http.Client{Timeout: 180 * time.Second},
		Syncer:        mautrix.NewDefaultSyncer(),
		Log:           zerolog.New(logrus.New().Out),
		Store:         dbSyncStore{},
	}

	_, err = client.Login(ctx, &mautrix.ReqLogin{
		Type: mautrix.AuthTypePassword,
		Identifier: mautrix.UserIdentifier{
			Type: mautrix.IdentifierTypeUser,
			User: config.C.Matrix.User.Localpart(),
		},
		Password:           config.C.Matrix.Password,
		DeviceID:           "matrix-meshtastic-bridge",
		StoreCredentials:   true,
		StoreHomeserverURL: true,
	})
	if err != nil {
		logrus.WithError(err).Fatal("failed to login to matrix")
	}

	return nil
}

func Run(ctx context.Context) {
	for {
		if err := client.Sync(); err != nil {
			logrus.WithError(err).Error("error syncing with matrix")
		}
	}
}

func Shutdown() {
	client.StopSync()
}

func SendMessage(ctx context.Context, text string) error {
	_, err := client.SendText(ctx, config.C.Matrix.Room, text)
	if err != nil {
		return err
	}

	return nil
}
