FROM alpine:3.18
WORKDIR /server
COPY out/checkout .
ENTRYPOINT ["/server/checkout"]
