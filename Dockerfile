FROM golang:latest

LABEL maintainer="Quique <hello@pragmaticreviews.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8080

RUN go build

CMD [ "swag-gin-demo" ]




