# syntax=docker/dockerfile:1
FROM golang:alpine as go-provider-builder

# Maintainer
MAINTAINER Ole Larssen <ole.larssen777@gmail.com>

ARG PORT

ENV PORT      $PORT

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    USER=appuser \
    UID=10001

RUN apk add --update --no-cache make git gcc build-base

# See https://stackoverflow.com/a/55757473/12429735RUN
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

#https://docs.docker.com/language/golang/build-images/
WORKDIR $GOPATH/src/provider/
COPY . .
RUN go build ./cmd/provider-service-server

# STEP 2 executable binary
FROM alpine:latest as provider

# redefine system variables in production image
ARG PORT

# set environment variables
ENV PORT     ${PORT}

# get GOPATH variable
FROM go-provider-builder
ENV GOPATH ${GOPATH}

WORKDIR /usr/local/bin/

# copy compiled binary and start the app
COPY --from=go-provider-builder .$GOPATH/src/provider/.env ./.env
COPY --from=go-provider-builder .$GOPATH/src/provider/provider-service-server ./provider-service-server
ENTRYPOINT ./provider-service-server --port=$PORT --host="0.0.0.0"
EXPOSE $PORT