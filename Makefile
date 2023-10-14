AWS_LAMBDA_FUNCTION_NAME=todo-list
ZIP_FILENAME=lambda-handler.zip

install:
	go install github.com/cosmtrek/air@latest

build:
	go build -C cmd/http -o ../../bootstrap

build-linux-arm64:
	GOOS=linux GOARCH=arm64 make build

dev:
	air

zip:
	zip $(ZIP_FILENAME) bootstrap

clean:
	rm -rf bootstrap $(ZIP_FILENAME)

deploy:
	aws s3 cp $(ZIP_FILENAME) s3://$(AWS_BUCKET)/ && \
	aws lambda update-function-code --function-name $(AWS_LAMBDA_FUNCTION_NAME) --s3-bucket $(AWS_BUCKET) --s3-key $(ZIP_FILENAME) >> /dev/null

test:
	go test -v ./...
