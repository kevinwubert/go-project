package util

import (
	"io/ioutil"
	"os"
)

// CreateDir creates a directory and returns an error
func CreateDir(name string) error {
	err := os.Mkdir(name, 0777)
	return err
}

// CreateFile creates a file with data and returns an error
func CreateFile(name string, data []byte) error {
	err := ioutil.WriteFile(name, data, 0644)
	return err
}
