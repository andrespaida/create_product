# Etapa de build
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# Etapa final (contenedor más liviano)
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

# Copiar archivo .env si lo usas localmente (opcional).
# COPY .env .

EXPOSE 4002

CMD ["./main"]