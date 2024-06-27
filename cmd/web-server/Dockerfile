FROM golang:1.22.1-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd/web-server/ ./web-server/
COPY ./pkg/ ./pkg/

WORKDIR /app/web-server
RUN go build -o web-server .

CMD ["./web-server"]
