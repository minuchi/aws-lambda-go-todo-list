ZIP_FILENAME=todo-list-lambda-handler.zip

install:
	go install github.com/cosmtrek/air@latest

build:
	GOOS=linux GOARCH=amd64 go build -C cmd/http -o ../../bootstrap

dev:
	air

zip:
	zip $(ZIP_FILENAME) bootstrap

clean:
	rm -rf bootstrap $(ZIP_FILENAME)

deploy:
	aws s3 cp $(ZIP_FILENAME) s3://$(AWS_BUCKET)/ && \
	aws lambda update-function-code --function-name todo-list --s3-bucket $(AWS_BUCKET) --s3-key $(ZIP_FILENAME)

test:
	go test -v ./...
