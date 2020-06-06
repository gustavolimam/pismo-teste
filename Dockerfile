FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    DB_CS="root:root@tcp(db:3306)/pismo" \
    PORT=":3003"

WORKDIR /build

COPY go.mod .
COPY go.sum .
COPY dbinit.sql /docker-entrypoint-initdb.d/dbinit.sql
RUN go mod download
RUN go get github.com/go-sql-driver/mysql

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

EXPOSE 3003

CMD ["/dist/main"]