package main

import "fmt"

func main() {
	/* const x = 3
	var y int16 = 12

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", x+y, x+y) */

	const x = int16(5.2 + 4.8)
	fmt.Printf("%T, %v\n", x, x)
}
