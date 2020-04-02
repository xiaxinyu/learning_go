package main

import (
	"strconv";
	"fmt"
)

func main() {
	s,_ := strconv.ParseInt("100", 10, 64)
	fmt.Printf("t=%T, s=%d",s, s)
}