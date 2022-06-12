package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {

	fileInfo, err := os.Stat("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Name: ", fileInfo.Name())
	fmt.Println("Size in Bytes: ", fileInfo.Size())
	fmt.Println("Permission: ", fileInfo.Mode())
	fmt.Println("Last Modified: ", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Println("System Info: ", fileInfo.Sys())

}
