FROM golang:alpine
WORKDIR /payment-service
COPY .. .
RUN go mod tidy
CMD ["go", "run", "cmd/main.go"]