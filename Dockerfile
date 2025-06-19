# syntax=docker/dockerfile:1
FROM golang:1.21

# Establecer directorio de trabajo
WORKDIR /app

# Copiar archivos de dependencia
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copiar el código restante
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto del microservicio
EXPOSE 4002

# Comando de ejecución
CMD ["./main"]