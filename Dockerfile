# Use the official image as a parent image.
FROM golang:latest as  bui
# Set the working directory in the image.
WORKDIR /app

# Copy the local package files to the container's workspace.
COPY . .

# Set up the Go environment and build the application.
ENV CGO_ENABLED=0 
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOPROXY="https://goproxy.io"
RUN go get -d -v

# Build the application.
RUN go build -o server .

# --------------------------------------
# START NEW STAGE
# --------------------------------------

# Use a slightly smaller base image from the Alpine repositories
# for a smaller final image size.
FROM alpine:latest  

# Refresh the package manager.
RUN apk --no-cache add ca-certificates

# Set the working directory.
WORKDIR /root/

# Copy the binary file from the first stage.
COPY --from=bui /app/server .

# Expose the app on port 8080.  
EXPOSE 8080 

# Run the application.
CMD ["./server"] 