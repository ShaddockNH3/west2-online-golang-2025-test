FROM golang:1.23.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o task4_tiktok_app .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/task4_tiktok_app .

EXPOSE 8080

CMD ["./task4_tiktok_app"]
