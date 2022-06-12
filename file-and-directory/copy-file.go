package main

import (
	"io"
	"log"
	"os"
)

func main() {

	orgFile, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer orgFile.Close()

	newFile, err := os.Create("new_example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesWritten, err := io.Copy(newFile, orgFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Copied %d bytes.", bytesWritten)

	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}

}
