FROM golang:1.17.0-alpine as builder

ENV LOG_LEVEL=info

WORKDIR /app

COPY . .
 
RUN go mod download
RUN go build -o oauth ./src
 
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/oauth .
 
# EXPOSE 8081

# ENTRYPOINT ["./items-api"]
