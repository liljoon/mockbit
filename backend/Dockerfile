FROM golang:tip-alpine AS builder

WORKDIR /app

COPY . .

COPY woongjin.cer /usr/local/share/ca-certificates/woongjin.cer
RUN cat woongjin.cer >> /etc/ssl/certs/ca-certificates.crt

RUN go mod download

RUN go build -o server .

CMD ["./server"]

#####

FROM alpine:latest

WORKDIR /dist

COPY --from=builder /app/server ./server

EXPOSE 8080

CMD ["./server"]
