FROM golang:1.22-alpine as svc-mutasi

RUN go install github.com/githubnemo/CompileDaemon@latest
RUN go install github.com/spf13/viper@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

ENTRYPOINT CompileDaemon --build="go build -o svc-mutasi" -command="./svc-mutasi" -build-dir=/app

