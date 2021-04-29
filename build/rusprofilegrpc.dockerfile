FROM alpine:latest as base

# prepare
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app

# restore packages
COPY ["go.mod", "go.sum", "./"]
RUN go get -d -v ./...

# build
COPY . .
RUN go install -v ./...

# run
FROM base as run
COPY --from=builder /go/bin/rusprofilegrpc /app
ENTRYPOINT ["./app"]
