package main         
import (
	"fmt"
) 

func main() {
	var s1, s2 string
	n, _ := fmt.Scan(&s1, &s2)
	fmt.Println(n, "개 입력")
	fmt.Println("s1: ", s1, " s2: ", s2)
}