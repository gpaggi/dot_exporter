FROM golang:1.16.5-alpine3.13 AS builder
ARG VERSION
WORKDIR /go/app

RUN apk add --no-cache bash git

COPY . /go/app/
RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o bin/dot_monitor -ldflags "-X github.com/gpaggi/dot_monitor/version.Version=$(VERSION) -s -w" .

FROM alpine:3.13
RUN adduser -s /sbin/nologin -D -H -u 1000 -g dot_monitor dot_monitor

USER dot_monitor

COPY --from=builder /go/app/bin/dot_monitor /usr/sbin/dot_monitor

CMD ["/usr/sbin/dot_monitor"]