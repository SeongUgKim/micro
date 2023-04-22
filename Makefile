build:
	go build -o bin/pricefetcher

client: build
	./bin/pricefetcher -mode=client

server: build
	./bin/pricefetcher -mode=server