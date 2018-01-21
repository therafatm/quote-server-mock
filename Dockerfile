FROM golang:latest

EXPOSE 8000

WORKDIR /go/src/quote_server_mock
ADD . /go/src/quote_server_mock



RUN go get github.com/pilu/fresh
RUN go get ./...


#RUN go build app.go

ENTRYPOINT ["fresh", "app.go"]
