# Build Stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/main.go

# Final Run Stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
# We don't copy .env because we will set variables in Render's dashboard
EXPOSE 8080
CMD ["./main"]