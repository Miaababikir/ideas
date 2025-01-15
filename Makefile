.DEFAULT_GOAL := run

.PHONY:run

build:
	go build -o bin/ideas cmd/main.go

run: build
	./bin/ideas

clean:
	go clean
