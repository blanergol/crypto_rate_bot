FROM golang:1.21.4-alpine3.18

RUN apk add build-base

ENV CGO_ENABLED=1 \
    GO111MODULE=on

WORKDIR /opt/app

COPY . .

RUN go build -tags musl -o /go/bin/service --ldflags "-extldflags -static -w -s" /opt/app/cmd/main.go

ENTRYPOINT ["/go/bin/service"]
