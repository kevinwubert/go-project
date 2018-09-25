all: build

build:
	go build -o bin/go-project main.go

clean:
	rm -rf bin