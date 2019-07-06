package main

func main() {
	i := 3
	switch i {
	case 2, 4, 6:
		println("짝수입니다.")
	case 1, 3, 5:
		println("홀수입니다.")
	}
}
