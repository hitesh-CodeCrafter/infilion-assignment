build:
	@go build -o bin/assignment

run: build
	@./bin/assignment