.PHONY: all build build-go-project build-process-templates install clean

all: build

build: build-process-templates run-process-templates build-go-project

build-go-project:
	go build -o bin/go-project main.go

build-process-templates:
	go build -o bin/process-templates scripts/processTemplates.go

run-process-templates:
	./bin/process-templates

install:
	go install

clean:
	rm -rf bin