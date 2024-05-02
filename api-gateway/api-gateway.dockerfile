#base go image
FROM golang:1.22-alpine3.19 AS builder

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY docs docs
COPY main.go .
COPY pkg pkg


RUN go build -o gateway ./cmd/api

FROM alpine:3.19

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/gateway .
COPY .env .

CMD ["sh","-c","echo $PORT && ./gateway"]





