package main

func main() {
	i := 2
	s := "hello"

	switch i {
	case 1:
		println(1)
	case 2:
		if s == "hello" {
			println("hello man")
			break
		}
		println("can't arrive")
	}
}
