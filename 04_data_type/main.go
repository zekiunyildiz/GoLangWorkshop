package main

import "fmt"

func main() {

	/* x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	fmt.Println(x + int(y)) */

	var x int8 = 127
	var y int16

	y = int16(x)

	fmt.Println(y)

}
