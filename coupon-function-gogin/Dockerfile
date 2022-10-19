# syntax=docker/dockerfile:1

FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

ENV GIN_MODE=release

RUN go build -o /coupon-action

FROM alpine:latest

RUN apk add --no-cache curl

RUN \
  export ARCH=amd64 && \
    curl -JLO "https://dl.filippo.io/mkcert/latest?for=linux/${ARCH}" && \
    chmod +x mkcert-v*-linux-${ARCH} && \
    cp mkcert-v*-linux-${ARCH} /usr/local/bin/mkcert && \
    mkdir /var/lib/caroot
ENV TRUST_STORES=system
ENV CAROOT=/var/lib/caroot
COPY rootCA.pem /var/lib/caroot
RUN mkcert -install

COPY --from=build /coupon-action /coupon-action

ENV PORT=80

CMD [ "/coupon-action" ]