package main

import (
	"log"
	"os"
)

func main() {

	orgPath := "example.txt"
	newPath := "./newFile/first.txt"

	err := os.Rename(orgPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
