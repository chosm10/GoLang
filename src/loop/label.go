package main

func main() {

Loop:
	for i := 10; i >= 5; i-- {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break Loop
			}
			println(i, j)
		}
	}

	println("Hello, world!")
}
