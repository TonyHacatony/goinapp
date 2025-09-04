# Этап сборки
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum и качаем зависимости
COPY go.mod ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем бинарь (статический)
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Финальный минимальный образ
FROM alpine:latest

WORKDIR /app

# Копируем бинарь из builder
COPY --from=builder /app/myapp .

# Делаем исполняемым
RUN chmod +x myapp

# Запуск приложения
CMD ["./myapp"]
