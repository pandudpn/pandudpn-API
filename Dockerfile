FROM golang:1.16-alpine3.13 AS builder

RUN apk update && apk add git

WORKDIR $GOPATH/src/pandudpn

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /go/bin/pandudpn cmd/main.go

FROM alpine:3.11

COPY --from=builder /go/bin/pandudpn /go/bin/pandudpn

RUN apk add --no-cache tzdata

ENTRYPOINT ["/go/bin/pandudpn"]