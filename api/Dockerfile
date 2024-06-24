ARG GO_VERSION=1.22

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go mod tidy
RUN go build -o ./app ./main.go

FROM golang:${GO_VERSION}-alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=builder /app .
COPY --from=builder /go/bin/air /usr/local/bin/
