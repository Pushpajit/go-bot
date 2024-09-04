# Step 1: Use an official Go image as the build environment
FROM golang:1.23-alpine AS builder

# Step 2: Set the Current Working Directory inside the container
WORKDIR /go-bot

# Step 3: Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Step 4: Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Step 5: Copy the source code into the container
COPY . .

# Step 6: Build the Go app
RUN go build -o gobot .

# Step 7: Use a smaller base image to reduce the final image size
FROM alpine:latest


# Step 9: Set the Current Working Directory inside the container for the final image
WORKDIR /root/

# Step 10: Copy the binary from the builder stage
COPY --from=builder /go-bot .

# Step 11: Expose the port that the app listens on (optional, change if necessary)
EXPOSE 8080

# Step 12: Command to run the executable
CMD ["./gobot"]
