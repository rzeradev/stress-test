FROM golang:1.21.3-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/

RUN go build -o /stress-test ./cmd/main.go

ENTRYPOINT ["/stress-test"]

