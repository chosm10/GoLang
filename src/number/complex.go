package main

func main() {
	var num1 complex64 = 1 + 2i
	var num2 complex64 = 4.2342 + 2.3527i

	var num3 complex64 = 1.e+3i
	var num4 complex64 = 7.2743e-10i

	var num5 complex128 = 1 + 2.3527i
	var num6 complex128 = 5.5322e-10 + .123456E+2i
	println(num1, num2, num3, num4, num5, num6)

	var r1 float64 = real(num6)
	var i1 float64 = imag(num6)

	var r2 float32 = real(num3)
	var i2 float32 = imag(num3)

	println(r1, i1)
	println(r2, i2)

	var num7 complex64 = complex(1, 2)
	print(num7)
}
