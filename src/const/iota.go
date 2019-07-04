package main

func main() {
	// const (
	// 	Sunday       = 0
	// 	Monday       = 1
	// 	Tuesday      = 2
	// 	Wednesday    = 3
	// 	Thursday     = 4
	// 	Friday       = 5
	// 	Saturday     = 6
	// 	numberOfDays = 7
	// )
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)
	println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, numberOfDays)

	const (
		a = 1 << iota
		b = 1 << iota
		c = 1 << iota
		d = 1 << iota
	)
	println(a, b, c, d)

	const (
		e = 30 * iota
		f = 30 * iota
		g = 30 * iota
		h = 30 * iota
	)
	println(e, f, g, h)

	const (
		bit0, mask0 = 1 << iota, 1<<iota - 1
		bit1, mask1
		_, _
		bit3, mask3
	)
	println(bit0, mask0)
	println(bit1, mask1)
	println(bit3, mask3)

	q := 1
	p := &q
	println(p)

}
