#!/bin/bash -e

PROTOC=./bin/protoc
INCLUDE_PATH="./include"

protoc_in_path=$(which protoc)
if [[ -x "${protoc_in_path}" ]]; then
  PROTOC=${protoc_in_path}
fi

$PROTOC \
  --go_out=plugins=grpc:. \
  -I=".:${INCLUDE_PATH}" \
  ./helloworld/*.proto

$PROTOC \
  --grpc-gateway_out=:. \
  --swagger_out=:. \
  -I=".:${INCLUDE_PATH}" \
  ./helloworld/*.proto
