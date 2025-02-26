FROM golang:1.22 AS builder

RUN apt-get update

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV ROOT_PATH=/go/src/app

RUN go install github.com/vektra/mockery/v2@v2.43.0
RUN go install github.com/swaggo/swag/cmd/swag@v1.7.8