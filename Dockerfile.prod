FROM golang:1.24.6-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o app ./cmd/main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
COPY .env .
CMD ["./app"]

# run command in terminal : sudo docker build -t docker-learn .
# to run the image : sudo docker run docker-learnt