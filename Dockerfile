# build
FROM golang:1.19.3-alpine as builder
RUN             apk add --no-cache git gcc musl-dev make
ENV             GO111MODULE=on
WORKDIR         /go/src/moul.io/graphman
COPY            . ./
RUN             make install

# minimalist runtime
FROM alpine:3.14.0
COPY            --from=builder /go/bin/pertify /bin/
