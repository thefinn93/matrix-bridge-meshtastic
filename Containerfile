FROM library/golang:1.25 AS build
ADD . /bridge
WORKDIR /bridge
RUN go build ./cmd/matrix-bridge-meshtastic

FROM gcr.io/distroless/base-debian12:nonroot
COPY --from=build /bridge/matrix-bridge-meshtastic /usr/bin/matrix-bridge-meshtastic
ENTRYPOINT ["/usr/bin/matrix-bridge-meshtastic"]
