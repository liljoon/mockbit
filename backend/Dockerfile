FROM golang:tip-alpine AS builder

WORKDIR /app

COPY woongjin.cer woongjin.cer
RUN cat woongjin.cer >> /etc/ssl/certs/ca-certificates.crt

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go build -o server .

CMD ["./server"]

#####

FROM alpine:latest

WORKDIR /dist

COPY --from=builder /app/server ./server

EXPOSE 8080

CMD ["./server"]
