FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go mod tidy

COPY . .

RUN go build -o teklif cmd/app/main.go

FROM debian:latest

WORKDIR /app

COPY --from=builder /app/teklif .

EXPOSE 8080

CMD ["./teklif"]
