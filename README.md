# dohproxy

##DNS over HTTPS proxy written in golang

This is a slight modification to satran/dohproxy to run the DNS-over-HTTPS proxy as a docker container/service.

To build, run the following command:
```
$ make
$ docker build -t dohproxy .
```

To run the container:
```
$ docker run -d --rm \
        --name dnsproxy \
        -e "DEBUG=<TRUE|FALSE>" \
        -e "SERVER_HOSTNAME=https://<dohprovider.url/path>" \
        -p 53:53/udp \
        --dns 1.1.1.1 \
        dohproxy
```

For the DOH provider URL, the default is `mozilla.cloudflare-dns.com/dns-query`. You can also use `https://cloudflare-dns.com/dns-query`.