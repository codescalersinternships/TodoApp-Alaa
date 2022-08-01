FROM golang:latest

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /swag-gin-demo

EXPOSE 8080

CMD [ "/swag-gin-demo" ]
