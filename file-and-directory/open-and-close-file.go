package main

import (
	"log"
	"os"
)

func main() {

	// Second Parameter: Control Behavior
	// Third Parameter: File Permission
	file, err := os.OpenFile("example.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

}
