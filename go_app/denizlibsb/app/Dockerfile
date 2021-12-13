FROM golang:alpine as builder

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /denizlibsbApp

EXPOSE 8080

CMD [ "/denizlibsbApp" ]