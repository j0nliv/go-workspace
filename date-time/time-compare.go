package main

import (
	"fmt"
	"time"
)

func main() {

	x := fmt.Println

	now := time.Now()
	x(now)

	x("-----------")
	x(now.Location())
	x(now.Day())
	x(now.Hour())
	x(now.Weekday())

	// Compare Date

	old := time.Date(1997, 4, 14, 10, 10, 0, 0, time.UTC)

	x(old.Before(now))
	x(old.After(now))

	diff := now.Sub(old)
	x(diff)
	x(diff.Hours())

}
