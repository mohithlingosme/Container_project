FROM golang:1.18 AS builder

WORKDIR /app

COPY app/go.mod app/go.sum ./
RUN go mod download

COPY app/main.go .
RUN go build -o main main.go

FROM ubuntu:18.04

COPY --from=builder /app/main /app/guestbook
COPY app/public/index.html /app/public/index.html
COPY app/public/script.js /app/public/script.js
COPY app/public/style.css /app/public/style.css
COPY app/public/jquery.min.js /app/public/jquery.min.js

WORKDIR /app

CMD ["./guestbook"]

EXPOSE 3000
