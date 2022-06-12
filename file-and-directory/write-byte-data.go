package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("example.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("New data!")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Wrote %d bytes.", bytesWritten)

}
