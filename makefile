all:
	CGO_ENABLED=0 go build -v
start:
	./zevola
clean:
	go fmt ./...