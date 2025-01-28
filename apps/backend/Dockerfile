FROM golang:1.23
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./src/cmd/main.go

# Use a minimal base image for the final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the build stage
COPY --from=0 /app/main .

# Expose the port your application runs on
EXPOSE 50051

CMD ["./main"]
