build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -C cmd/app -o ../../bootstrap

zip:
	make build-linux-amd64
	zip lambda-handler.zip bootstrap
	rm bootstrap

test:
	go test -v ./...
