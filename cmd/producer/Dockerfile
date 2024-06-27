FROM golang:1.22.1-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd/producer/ ./producer/
COPY ./pkg/ ./pkg/

WORKDIR /app/producer
RUN go build -o producer .

CMD ["./producer"]
