ARG PORT
ARG APP_NAME
ARG PORT
ARG NODE_ENV

# STEP 1 build executable binary
FROM golang:alpine as go-provider-builder

ARG PORT
ARG APP_NAME

ENV NODE_ENV $NODE_ENV
ENV PORT ${PORT}

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    USER=appuser \
    UID=10001

ENV GOPRIVATE=github.com

RUN apk add --update --no-cache bash git build-base gcc abuild binutils binutils-doc gcc-doc jq curl

ENV PACKAGES="\
  py-pip \
  jq \
"

RUN apk add --update $PACKAGES \
  && pip install yq \
  && rm /var/cache/apk/*

# See https://stackoverflow.com/a/55757473/12429735RUN
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
WORKDIR $GOPATH/src/provider/
COPY . .
RUN make build
ENTRYPOINT ./provider-service-server --port=$PORT --host="0.0.0.0"
EXPOSE $PORT
