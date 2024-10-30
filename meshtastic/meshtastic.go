package meshtastic

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"git.janky.solutions/finn/matrix-bridge-meshtastic/config"
	"git.janky.solutions/finn/matrix-bridge-meshtastic/meshtastic/protobufs"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

const (
	pathFromRadio = "/api/v1/fromradio"
	pathToRadio   = "/api/v1/toradio"
)

func Receive(ctx context.Context, ch chan *protobufs.FromRadio) {
	t := time.NewTicker(config.C.Meshtastic.PollingInterval)
	for {
		select {
		case <-t.C:
			fromRadio, err := getFromRadio(ctx)
			if err != nil {
				logrus.WithError(err).Error("error communicating with radio")
			}

			if fromRadio != nil && fromRadio.PayloadVariant != nil {
				ch <- fromRadio
			}

		case <-ctx.Done():
			return
		}
	}
}

func ShutdownReceiver() {
	// TODO: wait for receive loop to cleanly exit
}

func SendConfigInit(ctx context.Context) error {
	return sendToRadio(ctx, &protobufs.ToRadio{
		PayloadVariant: &protobufs.ToRadio_WantConfigId{
			WantConfigId: 4, // chosen by fair dice roll. guaranteed to be random. - actually though, we dont do anything special when startup is over so we dont care about this.
		},
	})
}

func getFromRadio(ctx context.Context) (*protobufs.FromRadio, error) {
	body, err := request(ctx, http.MethodGet, pathFromRadio, nil)
	if err != nil {
		return nil, err
	}

	fromRadio := protobufs.FromRadio{}
	err = proto.Unmarshal(body, &fromRadio)
	if err != nil {
		return nil, fmt.Errorf("error parsing response from radio: %v", err)
	}

	return &fromRadio, nil
}

func sendToRadio(ctx context.Context, req *protobufs.ToRadio) error {
	requestBody, err := proto.Marshal(req)
	if err != nil {
		return fmt.Errorf("error marshalling ToRadio payload: %v", err)
	}

	resp, err := request(ctx, http.MethodPut, pathToRadio, requestBody)
	if err != nil {
		return err
	}

	logrus.WithField("resp_body", string(resp)).Debug("completed request to radio")

	return nil
}

func request(ctx context.Context, method, path string, body []byte) ([]byte, error) {
	u := url.URL{
		Scheme: "http",
		Host:   config.C.Meshtastic.Address,
		Path:   path,
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error building request to radio with path %s: %v", path, err)
	}
	req.Header.Add("Accept", "application/x-protobuf")

	ctx, cancel := context.WithTimeout(ctx, config.C.Meshtastic.RequestTimeout)
	defer cancel()

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("error making request to radio with path %s: %v", path, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected %s from radio to path %s", resp.Status, path)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response from radio path %s: %v", path, err)
	}

	return responseBody, nil
}
