package main

import (
	"fmt"

	"github.com/kevinwubert/go-project/pkg/templates"
)

func main() {
	err := templates.ProcessTemplatesDir("asdf")
	if err != nil {
		fmt.Printf("Error in generating static templates: %v\n", err)
	}
}
