FROM golang:1.16-buster

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_USER app
ENV APP_DIR /go/src/app

RUN mkdir $APP_DIR

COPY . $APP_DIR

WORKDIR $APP_DIR

CMD go run .
