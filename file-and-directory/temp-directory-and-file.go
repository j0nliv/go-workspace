package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	tmpDir, err := ioutil.TempDir("", "tmp")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created a temp directory! Directory Name: ", tmpDir)

	tmpFile, err := ioutil.TempFile(tmpDir, "tmpFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created a temp file! File Name: ", tmpFile)

	err = tmpFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(tmpFile.Name())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s file is deleted!", tmpFile.Name())
	}

	err = os.Remove(tmpDir)
	if err != nil {
		log.Fatal(err)
	}

}
