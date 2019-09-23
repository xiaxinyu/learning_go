package main

import "os"

func main() {
	var (
		a int
		b bool
		str string
	)

	a= 1
	b = false
	str = "summer"

	println(a, b, str)

	var (
		HOME = os.Getenv("HOME")
		USER = os.Getenv("USER")
		GOROOT = os.Getenv("GOROOT")
	)

	println(HOME, USER, GOROOT)
}
