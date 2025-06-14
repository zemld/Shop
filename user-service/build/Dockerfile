FROM golang:alpine
WORKDIR /user-service
COPY .. .
RUN go mod tidy
CMD ["go", "run", "cmd/main.go"]