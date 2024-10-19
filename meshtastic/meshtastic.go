package meshtastic

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/config"
	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/meshtastic/protobufs"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

var shutdownRequested = false

func Receive() error {
	for !shutdownRequested {
		fromRadio, err := getFromRadio()
		if err != nil {
			logrus.WithError(err).Error("error communicating with radio")
		}

		if fromRadio != nil && fromRadio.PayloadVariant != nil {
			logrus.Debugf("fromRadio: %+v", fromRadio)
		}

		time.Sleep(time.Second)

	}
	return nil
}

func getFromRadio() (*protobufs.FromRadio, error) {
	u := url.URL{
		Scheme: "http",
		Host:   config.C.Meshtastic.Address,
		Path:   "/api/v1/fromradio",
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error building request to radio: %v", err)
	}
	req.Header.Add("Accept", "application/x-protobuf")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to radio: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected %s from radio", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response from radio: %v", err)
	}

	fromRadio := protobufs.FromRadio{}
	err = proto.Unmarshal(body, &fromRadio)
	if err != nil {
		return nil, fmt.Errorf("error parsing response from radio: %v", err)
	}

	return &fromRadio, nil
}
