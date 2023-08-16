install:
	go install github.com/cosmtrek/air@latest

build:
	GOOS=linux GOARCH=amd64 go build -C cmd/app -o ../../bootstrap

dev:
	air

zip:
	zip todo-list-lambda-handler.zip bootstrap

clean:
	rm -rf bootstrap todo-list-lambda-handler.zip

deploy:
	aws s3 cp todo-list-lambda-handler.zip s3://$(AWS_BUCKET)/

test:
	go test -v ./...
