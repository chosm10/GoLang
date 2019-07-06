package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Grade        int
	Gender       string
	SchoolNumber string
}

type Student struct {
	Name string
	Data Data
}

type School struct {
	Name     string
	Address  string
	Students []Student
}

func main() {
	doc := `
	{
		"Name": "Super University",
		"Address": "Republic Of Korea",
		"Students": [{
			"Name": "Minho",
			"Data": {
				"Grade": 2,
				"Gender": "Man",
				"SchoolNumber": "134912"
			}
		}]
	}
	`
	var school School
	json.Unmarshal([]byte(doc), &school)
	fmt.Println(school)
}
