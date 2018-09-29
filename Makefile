all: build

build:
	go build -o bin/go-project *.go

install:
	go install

process-templates:
	go build -o bin/process-templates scripts/processTemplates.go

clean:
	rm -rf bin