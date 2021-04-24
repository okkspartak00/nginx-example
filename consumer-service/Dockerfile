FROM golang:1.15-alpine3.13
WORKDIR app
RUN apk update \
    && apk add sqlite \
    && apk add socat \
    && apk add build-base
RUN which sqlite3
RUN which gcc
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o consumer-service
EXPOSE 8081
CMD ["./consumer-service"]