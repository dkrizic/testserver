all: generate build

generate:
	go generate ./...

build: generate
	go build -o testserver .

run: build
	./testserver

image:
	docker build -t testserver .
