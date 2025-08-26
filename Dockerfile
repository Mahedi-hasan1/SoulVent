FROM golang:1.24.6-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o soulvent ./cmd/main.go
CMD ["./soulvent"]
