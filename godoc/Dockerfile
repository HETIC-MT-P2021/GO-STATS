FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /doc
WORKDIR /doc

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go get -u golang.org/x/tools/...

# Copy the code into the container
COPY . .

# Export necessary port
EXPOSE 6060

# Command to run when starting the container
CMD ["godoc", "-http=:6060"]