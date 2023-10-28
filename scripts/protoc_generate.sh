#!/usr/bin/env bash

set -eux

script_dir="$(dirname $0)"
PROTO_PATH="./proto"

GO_OUT="proto.pb"

cd "${script_dir}/.."


# Generated golang *.pb.go
docker build -t twitter-protoc-go docker/protoc

find "$GO_OUT" -not -name "$GO_OUT" | xargs rm -rf

find "$PROTO_PATH" -type f -iregex ".*.proto" -print0 \
| xargs -0 docker run -v "$(pwd):/pb" -w /pb --rm twitter-protoc-go \
    protoc \
    --proto_path="$PROTO_PATH" \
    --go_out=plugins="grpc:$GO_OUT" \
    --go_opt=paths=source_relative