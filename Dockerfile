# Stage 1: Build the Go application
FROM golang:1.23-alpine3.20 AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

# Stage 2: Create a smaller, final image
FROM alpine:3.20.3
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 8080
CMD ["./main"]