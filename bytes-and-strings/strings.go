package main

import (
	"fmt"
	"strings"
)

func main() {

	builder := strings.Builder{}

	for i := 0; i < 10; i++ {
		builder.WriteString("Data ")
	}

	res := builder.String()
	fmt.Println(res)

}
