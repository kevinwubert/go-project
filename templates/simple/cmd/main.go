package main

import (
	"fmt"

	"github.com/kevinwubert/PROJECT_NAME/pkg/server"
)

func main() {
	err := server.Main()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
