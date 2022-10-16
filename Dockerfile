FROM golang:1.19-alpine

WORKDIR /var/app

COPY handlers.go .
COPY main.go .
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go build

CMD ["/var/app/main"]