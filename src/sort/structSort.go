package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name  string
	score float32
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].name < s[j].name
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ByScore struct {
	Students
}

func (s ByScore) Less(i, j int) bool {
	return s.Students[i].score < s.Students[j].score
}

func main() {
	students := Students{
		{"Minho", 73},
		{"Youngho", 92},
		{"Jimin", 86},
	}

	sort.Sort(students)
	fmt.Println(students)

	sort.Sort(sort.Reverse(ByScore{students}))
	fmt.Println(students)
}
