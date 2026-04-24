FROM alpine:3.17

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

RUN wget -O /bin/goose https://github.com/pressly/goose/releases/download/v3.14.0/goose_linux_x86_64 && \
  chmod +x /bin/goose

RUN apk add --no-cache dos2unix

WORKDIR /root

ADD chat-server/migrations/*.sql chat-server/migrations/
ADD auth/migrations/*.sql auth/migrations/
ADD migration.sh .
ADD .env .

RUN dos2unix .env migration.sh && chmod +x migration.sh

ENTRYPOINT ["bash", "migration.sh"]