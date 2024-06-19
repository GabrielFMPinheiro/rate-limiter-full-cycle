FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

RUN chmod 755 /app

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]