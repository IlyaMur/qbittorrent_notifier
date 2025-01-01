APP_NAME = notifier_bot

build:
	@go build -o $(APP_NAME)

run:
	@go run .

clean:
	@rm -f $(APP_NAME)

test:
	@go test -v ./...

.PHONY: build run clean test