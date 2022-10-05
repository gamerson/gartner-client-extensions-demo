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

COPY --from=build /coupon-action /coupon-action

ENV PORT=80

CMD [ "/coupon-action" ]