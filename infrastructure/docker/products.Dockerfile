FROM alpine:3.18
WORKDIR /server
COPY out/products .
ENTRYPOINT ["/server/products"]
