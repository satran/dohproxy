build:
	dep ensure -v
	CGO_ENABLED=0 go build -o dohproxy