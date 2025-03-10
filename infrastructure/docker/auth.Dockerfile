FROM alpine:3.18
WORKDIR /server
COPY out/auth .
ENTRYPOINT ["/server/auth"]
