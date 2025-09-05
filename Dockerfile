FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o myapp ./cmd/main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/myapp .
COPY --from=builder /app/config/config.json /app/config/config.json
EXPOSE 8080
CMD ["./myapp"]