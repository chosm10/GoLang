package main

func main() {
	i := 7

	switch {
	case i >= 5 && i < 10:
		println("5이상 10이하")
	case i >= 0 && i < 5:
		println("0이상 5이하")
	}
}
