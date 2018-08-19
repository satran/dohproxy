FROM alpine:3.7

RUN apk -U add ca-certificates

WORKDIR /var/srv/dns

COPY ./dohproxy .

CMD ["./dohproxy"]