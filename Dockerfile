# syntax=docker/dockerfile:1
FROM golang:1.24.2

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/server

EXPOSE 3000

CMD ["/app/main"]
