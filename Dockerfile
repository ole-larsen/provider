# syntax=docker/dockerfile:1

FROM golang:alpine as go-provider-builder

ARG PORT=3000
ARG NEXT_PUBLIC_GOOGLE_CLIENT_ID
ARG NEXT_PUBLIC_GOOGLE_CLIENT_SECRET
ARG NEXT_PUBLIC_GOOGLE_CLIENT_CALLBACK
ARG NEXT_PUBLIC_GOOGLE_CLIENT_REDIRECT

ARG NEXT_PUBLIC_YANDEX_CLIENT_ID
ARG NEXT_PUBLIC_YANDEX_CLIENT_SECRET
ARG NEXT_PUBLIC_YANDEX_CLIENT_CALLBACK
ARG NEXT_PUBLIC_YANDEX_CLIENT_REDIRECT

ARG NEXT_PUBLIC_TELEGRAM_BOT_KEY
ARG NEXT_PUBLIC_TELEGRAM_PUBLIC_KEY
ARG NEXT_PUBLIC_TELEGRAM_AUTH_URL

ARG DOMAIN
ARG NEXT_PUBLIC_URL
ENV PORT=$PORT
ENV NEXT_PUBLIC_GOOGLE_CLIENT_ID=${NEXT_PUBLIC_GOOGLE_CLIENT_ID}
ENV NEXT_PUBLIC_GOOGLE_CLIENT_SECRET=${NEXT_PUBLIC_GOOGLE_CLIENT_SECRET}
ENV NEXT_PUBLIC_GOOGLE_CLIENT_CALLBACK=${NEXT_PUBLIC_GOOGLE_CLIENT_CALLBACK}
ENV NEXT_PUBLIC_GOOGLE_CLIENT_REDIRECT=${NEXT_PUBLIC_GOOGLE_CLIENT_REDIRECT}

ENV NEXT_PUBLIC_YANDEX_CLIENT_ID=${NEXT_PUBLIC_YANDEX_CLIENT_ID}
ENV NEXT_PUBLIC_YANDEX_CLIENT_SECRET=${NEXT_PUBLIC_YANDEX_CLIENT_SECRET}
ENV NEXT_PUBLIC_YANDEX_CLIENT_CALLBACK=${NEXT_PUBLIC_YANDEX_CLIENT_CALLBACK}
ENV NEXT_PUBLIC_YANDEX_CLIENT_REDIRECT=${NEXT_PUBLIC_YANDEX_CLIENT_REDIRECT}
ENV NEXT_PUBLIC_VK_CLIENT_ID=${NEXT_PUBLIC_VK_CLIENT_ID}
ENV NEXT_PUBLIC_VK_CLIENT_SECRET=${NEXT_PUBLIC_VK_CLIENT_SECRET}
ENV NEXT_PUBLIC_VK_CLIENT_CALLBACK=${NEXT_PUBLIC_VK_CLIENT_CALLBACK}


ENV NEXT_PUBLIC_TELEGRAM_BOT_KEY=${NEXT_PUBLIC_TELEGRAM_BOT_KEY}
ENV NEXT_PUBLIC_TELEGRAM_PUBLIC_KEY=${NEXT_PUBLIC_TELEGRAM_PUBLIC_KEY}
ENV NEXT_PUBLIC_TELEGRAM_AUTH_URL=${NEXT_PUBLIC_TELEGRAM_AUTH_URL}

ENV DOMAIN=${DOMAIN}

ENV NEXT_PUBLIC_URL=${NEXT_PUBLIC_URL}

ENV USER=appuser \
    UID=10001

ENV GOPATH ${GOPATH}

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

RUN echo $GOPATH

RUN go mod download && go mod tidy

RUN go build -o bin/provider-service-server ./cmd/provider-service-server

# STEP 2 executable binary
FROM alpine:latest as provider


ARG PORT=3000
ARG NEXT_PUBLIC_GOOGLE_CLIENT_ID
ARG NEXT_PUBLIC_GOOGLE_CLIENT_SECRET
ARG NEXT_PUBLIC_GOOGLE_CLIENT_CALLBACK
ARG NEXT_PUBLIC_GOOGLE_CLIENT_REDIRECT

ARG NEXT_PUBLIC_YANDEX_CLIENT_ID
ARG NEXT_PUBLIC_YANDEX_CLIENT_SECRET
ARG NEXT_PUBLIC_YANDEX_CLIENT_CALLBACK
ARG NEXT_PUBLIC_YANDEX_CLIENT_REDIRECT

ARG NEXT_PUBLIC_TELEGRAM_BOT_KEY
ARG NEXT_PUBLIC_TELEGRAM_PUBLIC_KEY
ARG NEXT_PUBLIC_TELEGRAM_AUTH_URL

ARG DOMAIN
ARG NEXT_PUBLIC_URL

ENV PORT=$PORT
ENV NEXT_PUBLIC_GOOGLE_CLIENT_ID=${NEXT_PUBLIC_GOOGLE_CLIENT_ID}
ENV NEXT_PUBLIC_GOOGLE_CLIENT_SECRET=${NEXT_PUBLIC_GOOGLE_CLIENT_SECRET}
ENV NEXT_PUBLIC_GOOGLE_CLIENT_CALLBACK=${NEXT_PUBLIC_GOOGLE_CLIENT_CALLBACK}
ENV NEXT_PUBLIC_GOOGLE_CLIENT_REDIRECT=${NEXT_PUBLIC_GOOGLE_CLIENT_REDIRECT}

ENV NEXT_PUBLIC_YANDEX_CLIENT_ID=${NEXT_PUBLIC_YANDEX_CLIENT_ID}
ENV NEXT_PUBLIC_YANDEX_CLIENT_SECRET=${NEXT_PUBLIC_YANDEX_CLIENT_SECRET}
ENV NEXT_PUBLIC_YANDEX_CLIENT_CALLBACK=${NEXT_PUBLIC_YANDEX_CLIENT_CALLBACK}
ENV NEXT_PUBLIC_YANDEX_CLIENT_REDIRECT=${NEXT_PUBLIC_YANDEX_CLIENT_REDIRECT}
ENV NEXT_PUBLIC_VK_CLIENT_ID=${NEXT_PUBLIC_VK_CLIENT_ID}
ENV NEXT_PUBLIC_VK_CLIENT_SECRET=${NEXT_PUBLIC_VK_CLIENT_SECRET}
ENV NEXT_PUBLIC_VK_CLIENT_CALLBACK=${NEXT_PUBLIC_VK_CLIENT_CALLBACK}

ENV NEXT_PUBLIC_TELEGRAM_BOT_KEY=${NEXT_PUBLIC_TELEGRAM_BOT_KEY}
ENV NEXT_PUBLIC_TELEGRAM_PUBLIC_KEY=${NEXT_PUBLIC_TELEGRAM_PUBLIC_KEY}
ENV NEXT_PUBLIC_TELEGRAM_AUTH_URL=${NEXT_PUBLIC_TELEGRAM_AUTH_URL}
ENV DOMAIN=${DOMAIN}
ENV NEXT_PUBLIC_URL=${NEXT_PUBLIC_URL}
ENV GOPATH=$GOPATH

RUN echo $GOPATH
WORKDIR /usr/local/bin/

COPY --from=go-provider-builder /go/src/bin/provider-service-server ./provider-service-server

RUN echo app is running on port > $PORT

# EXPOSE $PORT
ENTRYPOINT ./provider-service-server --port=$PORT --host="0.0.0.0"