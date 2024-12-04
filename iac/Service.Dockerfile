# Stage 1: Build the Go application
FROM golang:1.22.2-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY service /app/service/
RUN go build -o app service/tiny-url-implementation/router/router.go

# Stage 2: Create a minimal runtime image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]