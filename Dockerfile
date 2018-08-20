FROM golang AS golang
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /doh .
# We need a regular DNS server to resolve the HTTPS api. It's a chicken egg
# problem.
RUN echo "8.8.8.8" > /resolv.conf

FROM scratch
COPY --from=golang /doh /
COPY --from=golang /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=golang /resolv.conf /etc/resolv.conf
EXPOSE 53
ENTRYPOINT ["/doh"]
CMD ["-host", "0.0.0.0", "-port", "53"]
