package main

func main() {
	for i := 10; i >= 5; i-- {
		if i <= 7 {
			println("멈춤")
			break
		}
		println(i)
	}
}
