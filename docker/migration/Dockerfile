FROM alpine:3.16.0

RUN apk update && \
    apk upgrade && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.5.3/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /migrations