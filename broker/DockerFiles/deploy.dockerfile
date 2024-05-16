FROM golang:alpine as builder

RUN apk update

WORKDIR /app

ADD . /app

RUN cd ./main && \
    go build ./main.go && \
    chmod 755 ./main

EXPOSE 8000

CMD ["./main/main"]

