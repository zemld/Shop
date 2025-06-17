FROM golang:alpine
WORKDIR /order-service
COPY .. .
RUN go mod tidy
CMD ["go", "run", "cmd/main.go"]