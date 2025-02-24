FROM golang:1.19-alpine3.18 AS builder
WORKDIR /build
COPY . .
RUN go build -o server cmd/main.go

FROM alpine:3.18
WORKDIR /server
COPY --from=builder /build/server .
ENTRYPOINT ["/server/server"]
