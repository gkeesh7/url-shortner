# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

RUN apk update
RUN apk add git

ENV PKG_NAME=url-shortner
ENV PKG_PATH=$GOPATH/src/$PKG_NAME
WORKDIR $PKG_PATH


COPY glide.yaml glide.lock $PKG_PATH/
RUN wget https://github.com/Masterminds/glide/releases/download/v0.13.3/glide-v0.13.3-linux-386.tar.gz
RUN tar xzvf glide-v0.13.3-linux-386.tar.gz
RUN linux-386/glide install

COPY . $PKG_PATH
RUN go env -w GO111MODULE=off

WORKDIR $PKG_PATH
EXPOSE 8080
CMD ["go", "run", "main.go"]



