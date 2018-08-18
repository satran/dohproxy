FROM alpine:3.7

WORKDIR /var/srv/dns

COPY ./dohproxy .