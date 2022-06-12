package main

import (
	"log"
	"os"
)

func main() {

	err := os.Remove("example.txt")
	if err != nil {
		log.Fatal(err)
	}

}
