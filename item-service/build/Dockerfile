FROM golang:alpine
WORKDIR /item-service
COPY .. .
RUN go mod tidy
CMD ["go", "run", "cmd/main.go"]