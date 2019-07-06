package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	doc := `
	{
		"name" : "Maria",
		"age" : 10
	}
	`

	var data map[string]interface{}
	json.Unmarshal([]byte(doc), &data)
	fmt.Println(data["name"], data["age"])
}
