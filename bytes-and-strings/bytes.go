package main

import (
	"bytes"
	"fmt"
)

func main() {

	a := "qwe"
	b := "xyz"
	c := a + b

	fmt.Println(c)

	var x bytes.Buffer
	for i := 0; i < 10; i++ {
		x.WriteString(generateStr())
	}
	fmt.Println(x.String())

}

func generateStr() string {
	return "$xyz-12345"
}
