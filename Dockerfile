# Build stage
FROM golang:1.17 AS builder

ARG VERSION=dev

WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main -ldflags=-X=main.version=${VERSION} main.go

# Final stage
FROM ubuntu:20.04

# Copy the datasets.json file from your local directory to the image
COPY datasets.json /go/src/app/datasets.json

COPY --from=builder /go/src/app/main /go/bin/main
ENV PATH="/go/bin:${PATH}"
CMD ["main"]
