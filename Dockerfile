FROM golang:alpine

ARG SOURCE_COMMIT

ADD . /go/src/github.com/baltuonis/prometheus-bigbluebutton-exporter
WORKDIR /go/src/github.com/baltuonis/prometheus-bigbluebutton-exporter

RUN DATE=$(date -u '+%Y-%m-%d-%H%M UTC'); \
    go install -ldflags="-X 'main.Version=${SOURCE_COMMIT}' -X 'main.BuildTime=${DATE}'" ./...

ENTRYPOINT  [ "/go/bin/prometheus-bigbluebutton-exporter" ]
EXPOSE      9688
