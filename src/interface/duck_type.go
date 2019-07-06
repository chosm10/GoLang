package main

import "fmt"

type Duck struct {
}
type Person struct {
}

func (d Duck) quack() {
	fmt.Println("오리는 꽥꽥")
}
func (d Duck) feathers() {
	fmt.Println("오리는 깃털을 푸드덕푸드덕!")
}
func (p Person) quack() {
	fmt.Println("꽥꽥 소리내는 오리를 바라봅니다.")
}
func (p Person) feathers() {
	fmt.Println("빠진 깃털을 쳐다 봅니다.")
}
func intheforest(q Quacker) {
	q.quack()
	q.feathers()
}

type Quacker interface {
	quack()
	feathers()
}

func main() {
	var p Person
	var d Duck

	intheforest(p)
	intheforest(d)

	if value, ok := interface{}(p).(Quacker); ok {
		fmt.Println(value, ok)
	}
	if value, ok := interface{}(d).(Quacker); ok {
		fmt.Println(value, ok)
	}
	var i int
	if value, ok := interface{}(i).(Quacker); ok {
		fmt.Println(value, ok)
	} else {
		fmt.Println(value, ok)
	}
}
