# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

RUN apk update
RUN apk add git

ENV PKG_NAME=url-shortner
ENV PKG_PATH=$GOPATH/src/$PKG_NAME
WORKDIR $PKG_PATH

COPY . $PKG_PATH
RUN go mod download
RUN go build main.go

WORKDIR $PKG_PATH
EXPOSE 8080
CMD ["go", "run", "main.go"]
