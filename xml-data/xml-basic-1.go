package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	file, err := os.Open("notes.xml")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	v := ObjectNotes{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Println(v)

}

type ObjectNotes struct {
	XMLName xml.Name `xml:"notes"`
	Version string   `xml:"version,attr"`
	Notes   []note   `xml:note`
	Body    string   `xml:",innerxml"`
}

type note struct {
	XMLName xml.Name `xml:"note"`
	To      string   `xml:to`
	From    string   `xml:from`
	Heading string   `xml:heading`
	Body    string   `xml:body`
}
