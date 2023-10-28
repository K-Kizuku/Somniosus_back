FROM golang:1.21-alpine

ENV PROTOC_VERSION=3.24.4 \
    GO111MODULE=on \
    PATH="/usr/local/bin:/usr/bin:${PATH}"

WORKDIR /pb

RUN apk add --no-cache git protobuf-dev protoc \
    && go install github.com/golang/protobuf/protoc-gen-go@v1.5.3