package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("example.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Permission required!")
		}
	}
	file.Close()

}
