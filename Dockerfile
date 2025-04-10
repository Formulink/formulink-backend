FROM golang:1.24-alpine3.21 as builder


WORKDIR /app

COPY ../ .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /service ./cmd/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /service ./

EXPOSE 8080

CMD ["./service"]