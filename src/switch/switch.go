package main

func main() {
	i := 7

	switch i {
	case 0:
		println(0)
	case 1:
		println(1)
	case 2:
		println(2)
	case 3:
		println(3)
	case 4:
		println(4)
	default:
		println(-1)
	}

	s := "world"

	switch s {
	case "hello":
		println("Hello")
	case "world":
		println("world")
	default:
		println("Hell")
	}
}
