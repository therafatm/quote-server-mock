FROM golang:latest

EXPOSE 8000


ADD . /go/src/quote_server_mock
WORKDIR /go/src/quote_server_mock

RUN go get github.com/pilu/fresh
RUN go get ./...

ENTRYPOINT ["fresh", "app.go"]
