FROM golang:1.23.0-alpine as monitor-builder

RUN apk --no-cache add gcc bash make musl-dev

WORKDIR /app

# dependency installing
COPY go.mod go.sum ./
RUN go mod download

# build project
COPY . .
RUN go build -o monitor-service cmd/main.go