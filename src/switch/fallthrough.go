package main

func main() {

	i := 3

	switch i {
	case 4:
		println("4이상")
		fallthrough
	case 3:
		println("3이상")
		fallthrough
	case 2:
		println("2이상")
		fallthrough
	case 1:
		println("1이상")
		fallthrough
	case 0:
		println("0이상")
	}
}
