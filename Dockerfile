FROM golang:latest

EXPOSE 8000

COPY . /go/src/quote-server-mock
WORKDIR /go/src/quote-server-mock

RUN go get github.com/pilu/fresh
RUN go get -d ./...
