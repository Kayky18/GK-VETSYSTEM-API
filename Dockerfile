FROM golang:1.23 AS api

WORKDIR /usr/src/

COPY . .

EXPOSE 8080




