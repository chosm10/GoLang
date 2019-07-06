package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := make(map[string]interface{})

	data["Name"] = "Maria"
	data["age"] = 10

	doc1, _ := json.Marshal(data)
	doc2, _ := json.MarshalIndent(data, "", ",")
	fmt.Println(string(doc1))
	fmt.Println(string(doc2))
}
