# Stage 1: builder
FROM golang:1.24 AS builder
 
WORKDIR /app
 
COPY go.mod go.sum ./
RUN go mod download
 
COPY . .
COPY ./config/config.yml /app/config/config.yml
 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service ./cmd/main.go
 
# Stage 2: minimal runtime
FROM alpine:3.21
 
WORKDIR /app
 
COPY --from=builder /app/service .
COPY ./config/config.yml /app/config/config.yml
 
EXPOSE 8081
 
CMD ["./service"]
