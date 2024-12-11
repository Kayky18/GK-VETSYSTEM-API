FROM golang:1.23 AS api

WORKDIR /usr/src/

COPY . .

go 

RUN go run main.go

