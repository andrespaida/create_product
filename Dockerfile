FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Compila el binario correctamente
RUN go build -o create_product .

# Expone el puerto usado en tu .env
EXPOSE 4002

# Ejecuta el binario al arrancar el contenedor
CMD ["./create_product"]