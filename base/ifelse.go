package main

import (
	"fmt"
	"runtime"
)

func main() {
	bool1 := true
	if bool1 {
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}

	check()
}

var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

func check() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
	fmt.Println(prompt)
}
