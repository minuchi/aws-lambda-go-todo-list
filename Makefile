build:
	GOOS=linux GOARCH=amd64 go build -C cmd/app -o ../../bootstrap

zip:
	zip lambda-handler.zip bootstrap

clean:
	rm -rf bootstrap lambda-handler.zip

deploy:
	aws lambda update-function-code --function-name todo-list --zip-file fileb://./lambda-handler.zip

test:
	go test -v ./...
