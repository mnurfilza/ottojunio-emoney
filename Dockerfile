FROM golang:1.19.2-alpine3.16 as builder

RUN apk update && apk add --no-cache git

RUN mkdir /app

WORKDIR /app


COPY go.mod .
COPY go.sum .

COPY . .

ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait


RUN go mod tidy

RUN go build main.go



CMD /wait && ./main
