FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    DB_CS="root:root@tcp(127.0.0.1:33306)/pismo" \
    PORT=":3003"

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
COPY dbinit.sql /docker-entrypoint-initdb.d/dbinit.sql
RUN go mod download
RUN go get github.com/go-sql-driver/mysql

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Export necessary port
EXPOSE 3003

# Command to run when starting the container
CMD ["/dist/main"]