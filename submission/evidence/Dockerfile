FROM golang:1.18 AS builder

WORKDIR /app

COPY app/main.go .

RUN go mod init guestbook
RUN go mod edit '-require=github.com/codegangsta/negroni@v1.0.0' '-require=github.com/gorilla/mux@v1.8.0' '-require=github.com/xyproto/simpleredis/v2@v2.9.0'
RUN go mod tidy
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
