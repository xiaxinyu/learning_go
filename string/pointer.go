package main

import(
	"fmt"
)

func main() {
	//print pointer
	var i int = 100
	fmt.Println("i =",i, ", &i =", &i)

	//transter pointer to value
	var p *int = &i
	var t int = *p
	fmt.Println("p =",p, ", t =", t)
}