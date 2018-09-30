.PHONY: all build build-go-project build-process-templates install clean

all: build

build: build-go-project build-process-templates

build-go-project:
	go build -o bin/go-project main.go

build-process-templates:
	go build -o bin/process-templates scripts/processTemplates.go

install:
	go install

clean:
	rm -rf bin