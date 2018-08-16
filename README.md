# dohproxy

DNS over HTTPS proxy written in golang

I got interested in DNS over HTTPS after Firefox started supporting it in its latest release. I looked around to understand how it worked. Most of the implementations were too complex and did a lot of things. I read the RFC[1] and realised it was very trivial. So I tried my hand at implementing a proxy. This is just a proof of concept.

To install it you can use:
```
go get github.com/satran/dohproxy
```
This assumes you have installed go.


To run it use:
```
dohproxy
```
This will start the proxy on `5353` port.

You can resolve addresses using:
```
dig @127.0.0.1 -p 5353 redhat.com
```


[1] https://tools.ietf.org/html/draft-ietf-doh-dns-over-https-13
