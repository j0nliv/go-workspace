package main

import (
	"fmt"

	"example.com/sum"
)

func main() {
	total := sum.Total(3,5)
	fmt.Println(total)
}