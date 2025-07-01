FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# 🔧 Desactiva VCS stamping para evitar el error
RUN go build -buildvcs=false -o main .

EXPOSE 4002

CMD ["./main"]