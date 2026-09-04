FROM golang:1.27-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY . .

RUN go build -o calculator-server ./cmd/server

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/calculator-server .

EXPOSE 8080

CMD ["./calculator-server"]