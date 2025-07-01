FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# ✅ Evita el error VCS y compila durante build, no en CMD
RUN go build -buildvcs=false -o create_product .

EXPOSE 4002

# ✅ Ejecuta el binario directamente
CMD ["./create_product"]