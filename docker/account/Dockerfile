FROM golang:1.21 AS builder


WORKDIR /go/src/github.com/K-Kizuku/Somniosus_back

COPY go.* .
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN make build-account

FROM alpine:3.18 AS app

COPY --from=builder /go/src/github.com/K-Kizuku/Somniosus_back/bin/account /usr/local/bin/account

RUN apk add --no-cache ca-certificates

CMD ["account"]