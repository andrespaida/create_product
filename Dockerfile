FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -buildvcs=false -o create_product .

EXPOSE 4002

CMD ["./create_product"]