FROM golang:1.24.2-alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o involta-test /build/cmd/main.go

FROM alpine:3.19

WORKDIR /app

# Копируем бинарник из builder-этапа
COPY --from=builder /build/involta-test .

# Экспортируем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./involta-test"]