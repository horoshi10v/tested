# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o /app/app ./cmd/server

# Final stage
FROM golang:1.23-alpine
WORKDIR /app
COPY --from=builder /app/app /app/app
COPY . .
RUN chmod +x /app/app
EXPOSE 8081
CMD ["/app/app"]
