FROM golang:alpine AS build

WORKDIR /build

# Resolve dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build go binary
COPY . .
RUN go build -o shortdesc *.go

############################
FROM alpine:latest AS server

WORKDIR /server

# Copy over binary and config
COPY --from=build /build/shortdesc .
COPY --from=build /build/env.yaml .

CMD ["./shortdesc"]
