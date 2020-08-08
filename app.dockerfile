FROM golang:1.14-alpine

RUN apk add curl git

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/github.com/barokurniawan/websocket

EXPOSE 3001

COPY . /go/src/github.com/barokurniawan/websocket

RUN dep ensure

CMD ["sh", "-c", "go run main.go"]
