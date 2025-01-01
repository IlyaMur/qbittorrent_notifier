APP_NAME = notifier_bot

build:
	@go build -o $(APP_NAME)

run:
	@go run .

clean:
	@rm -f $(APP_NAME)

.PHONY: build run clean