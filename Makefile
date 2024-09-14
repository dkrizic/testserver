all: generate build test

generate:
	go generate ./...

build: generate
	go build -o testserver .

run: build
	./testserver

image:
	docker build -t testserver .

diagram:
	graphqlviz graph/schema.graphqls | dot -Tpng -o graph.png

clean:
	rm -f testserver

test:
	go test ./...
