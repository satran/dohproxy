# dohproxy

## Golang DNS-over-HTTPS proxy as a docker container/service.

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

The `--dns` argument is required for the initial name resolution of the DOH provider hostname. Here we are using Cloudflare's `1.1.1.1`, but any other secure DNS provider can be used (i.e. Quad9 - `9.9.9.9`).

Once the container is running, all you need to do is modify your DNS settings to use only `127.0.0.1` for name resolution.
