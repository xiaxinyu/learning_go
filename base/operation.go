package main

import (
	"fmt"
)

func main() {
	var a int
	var b int32
	a = 15
	a = a + a
	b = b + 5
	println(a, b)

	var n int16 = 34
	var m int32
	// compiler error: cannot use n (type int16) as type int32 in assignment
	//m = n
	m = int32(n)

	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)
}
