package main

func main() {
	var b1 bool = true
	var b2 bool = false

	println(b1, b2)
	println(true && true)   //t
	println(true && false)  //f
	println(false && false) //f
	println(true || true)   //t
	println(true || false)  //t
	println(false || false) //f
	println(!true)          //f
	println(!false)         //t

	var num1 int = 3
	var num2 int = 10

	println(num1 > num2)  //f
	println(num1 < num2)  //t
	println(num1 != num2) //t
	println(num1 >= num2) //f
	println(num1 <= num2) //t
}
