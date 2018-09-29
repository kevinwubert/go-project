all: build

build:
	go build -o bin/go-project main.go

install:
	go install

clean:
	rm -rf bin