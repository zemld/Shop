FROM golang:alpine
WORKDIR /admin-service
COPY .. .
RUN go mod tidy
CMD ["go", "run", "cmd/main.go"]