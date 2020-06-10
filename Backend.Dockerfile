FROM golang:latest

WORKDIR /usr/src/app

COPY . /usr/src/app

RUN go build server.go

CMD ["./server"]