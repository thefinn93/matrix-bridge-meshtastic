#!/bin/bash
set -exuo pipefail

for src in protobufs-src/meshtastic/*.proto protobufs-src/nanopb.proto; do
    dest="protobufs/$(basename "${src}")"
    cp -v "${src}" "${dest}"
    sed -i 's#option go_package =.*#option go_package = "git.janky.solutions/finn/matrix-bridge-meshtastic/meshtastic/protobufs";#' "${dest}"
    sed -i 's#import "meshtastic/#import "protobufs/#' "${dest}"
    sed -i 's#import "nanopb.proto#import "protobufs/nanopb.proto#' "${dest}"
done

protoc --go_out=. --go_opt=paths=source_relative --go_opt=Mprotobufs/nanopb.proto=git.janky.solutions/finn/matrix-bridge-meshtastic/meshtastic/protobufs protobufs/*.proto
