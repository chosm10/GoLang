package main

func main() {
	for i, j := 0, 0; i < 10; i, j = i+1, j+2 {
		println(i, j)
	}

	// 증감식에서 병렬 할당 안하고 따로 하면 에러
	// for i, j := 0, 0; i < 10; i++, j += 2 {
	// 	println(i, j)
	// }
}
