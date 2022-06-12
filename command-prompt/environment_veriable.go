package main

import (
	"fmt"
	"os"
)

func main() {
	uName := os.Getenv("USERNAME")
	uDomain := os.Getenv("DOMAIN")
	processorArch := os.Getenv("PROCESSOR_ARCHITECTURE")
	processorId := os.Getenv("PROCESSOR_IDENTIFIER")
	processorLevel := os.Getenv("PROCESSOR_LEVEL")
	goRoot := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")
	homePath := os.Getenv("HOMEPATH")

	fmt.Println("uname :" + uName)
	fmt.Println("domain :" + uDomain)
	fmt.Println("arch :" + processorArch)
	fmt.Println("id :" + processorId)
	fmt.Println("level :" + processorLevel)
	fmt.Println("root :" + goRoot)
	fmt.Println("go path :" + goPath)
	fmt.Println("home path :" + homePath)
}
