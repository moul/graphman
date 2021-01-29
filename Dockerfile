# build
FROM golang:1.15.7-alpine as builder
RUN             apk add --no-cache git gcc musl-dev make
ENV             GO111MODULE=on
WORKDIR         /go/src/moul.io/graphman
COPY            . ./
RUN             make install

# minimalist runtime
FROM alpine:3.13.1
COPY            --from=builder /go/bin/pertify /bin/
