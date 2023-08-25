FROM golang:1.20.4-alpine
WORKDIR /coffeebackend

COPY . .

RUN go mod download
RUN go build -v -o /coffeebackend/backend ./cmd/main.go

EXPOSE 8080

ENTRYPOINT [ "/coffeebackend/backend" ] 