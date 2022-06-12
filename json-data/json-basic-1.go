package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonData := `	
		{
			"data": {
				"object":"card",
				"id":"card_123456",
				"first":"john",
				"last":"morrison",
				"balance":"45.2"
			}	
		}
	`

	var c map[string]map[string]interface{}

	if err := json.Unmarshal([]byte(jsonData), &c); err != nil {
		panic(err)
	}

	fmt.Println(c)

	fmt.Println("-----")

	r, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(r))

}
