# syntax=docker/dockerfile:1

FROM golang:alpine as go-provider-builder

ARG PORT=3000
ARG NEXT_PUBLIC_GOOGLE_CLIENT_ID
ARG NEXT_PUBLIC_GOOGLE_CLIENT_SECRET
ARG NEXT_PUBLIC_GOOGLE_CLIENT_CALLBACK

ENV PORT=$PORT
ENV NEXT_PUBLIC_GOOGLE_CLIENT_ID=${NEXT_PUBLIC_GOOGLE_CLIENT_ID}
ENV NEXT_PUBLIC_GOOGLE_CLIENT_SECRET=${NEXT_PUBLIC_GOOGLE_CLIENT_SECRET}
ENV NEXT_PUBLIC_GOOGLE_CLIENT_CALLBACK=${NEXT_PUBLIC_GOOGLE_CLIENT_CALLBACK}

RUN echo $PORT > used port

ENV USER=appuser \
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
RUN mkdir -p $GOPATH/src/github.com/ole-larsen/provider
WORKDIR $GOPATH/src/

COPY . .

RUN go mod download

RUN go build -o bin/provider-service-server ./cmd/provider-service-server

# STEP 2 executable binary
FROM alpine:latest as provider

# get GOPATH variable
FROM go-provider-builder

ARG PORT=3000
ARG NEXT_PUBLIC_GOOGLE_CLIENT_ID
ARG NEXT_PUBLIC_GOOGLE_CLIENT_SECRET
ARG NEXT_PUBLIC_GOOGLE_CLIENT_CALLBACK

ENV PORT=$PORT
ENV NEXT_PUBLIC_GOOGLE_CLIENT_ID=${NEXT_PUBLIC_GOOGLE_CLIENT_ID}
ENV NEXT_PUBLIC_GOOGLE_CLIENT_SECRET=${NEXT_PUBLIC_GOOGLE_CLIENT_SECRET}
ENV NEXT_PUBLIC_GOOGLE_CLIENT_CALLBACK=${NEXT_PUBLIC_GOOGLE_CLIENT_CALLBACK}

ENV GOPATH ${GOPATH}

WORKDIR /usr/local/bin/

# copy compiled binary and start the app

COPY --from=go-provider-builder .$GOPATH/src/bin/provider-service-server ./provider-service-server

RUN echo app is running on port > $PORT

EXPOSE $PORT
ENTRYPOINT ./provider-service-server --port=$PORT --host="0.0.0.0"