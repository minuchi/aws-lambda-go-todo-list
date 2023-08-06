build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -C cmd/app -o ../../app

zip:
	make build-linux-amd64
	zip lambda-handler.zip app
	rm app
