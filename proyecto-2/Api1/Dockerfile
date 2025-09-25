# Etapa de compilación
FROM golang:1.24.4 AS builder
WORKDIR /app

# Copiar todo (incluye main.go y go.mod si existe)
COPY . .

# Compilar binario
RUN go build -o app

# Etapa final (imagen ligera)
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8081
CMD ["./app"]
