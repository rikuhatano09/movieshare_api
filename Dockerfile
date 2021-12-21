FROM golang:1.17-alpine3.13

WORKDIR /go/src

# Install Go modules.
COPY ./movieshare_api/go.mod ./
RUN go mod download
