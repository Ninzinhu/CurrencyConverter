# Estágio de construção (Go)
FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Estágio de execução
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY ./web ./web
EXPOSE 8080
CMD ["./server"]