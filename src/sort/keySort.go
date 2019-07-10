package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name  string
	score float32
}

type By func(s1, s2 *Student) bool

type studentSorter struct {
	students []Student
	by       func(s1, s2 *Student) bool
}

func (by By) Sort(students []Student) {
	sorter := &studentSorter{
		students: students,
		by:       by,
	}
	sort.Sort(sorter)
}

func (s *studentSorter) Len() int {
	return len(s.students)
}

func (s *studentSorter) Less(i, j int) bool {
	return s.by(&s.students[i], &s.students[j])
}

func (s *studentSorter) Swap(i, j int) {
	s.students[i], s.students[j] = s.students[j], s.students[i]
}

func main() {
	students := []Student{
		{"Maria", 78.2},
		{"Dongho", 52.3},
		{"Bongdu", 92.8},
	}

	name := func(s1, s2 *Student) bool {
		return s1.name < s2.name
	}

	score := func(s1, s2 *Student) bool {
		return s1.score < s2.score
	}

	reverseName := func(s1, s2 *Student) bool {
		return !name(s1, s2)
	}

	By(name).Sort(students)
	fmt.Println(students)

	By(score).Sort(students)
	fmt.Println(students)

	By(reverseName).Sort(students)
	fmt.Println(students)
}
