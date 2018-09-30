.PHONY: all build build-go-project build-process-templates run-process-templates fmt install clean

all: build

build: build-process-templates run-process-templates fmt build-go-project

build-go-project:
	go build -o bin/go-project main.go

build-process-templates:
	go build -o bin/process-templates scripts/processTemplates.go

run-process-templates:
	./bin/process-templates

fmt:
	go fmt

install:
	go install

clean:
	rm -rf bin