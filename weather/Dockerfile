FROM golang:1.23.5-alpine

WORKDIR /app

# Install required dependencies
RUN apk add --no-cache gcc musl-dev

# Copy go mod and sum files
COPY go.mod go.sum ./../

COPY .env ./../

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./weather

# Expose the port
EXPOSE 3007

# Run the application
CMD ["./main"]
