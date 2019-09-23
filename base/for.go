package main

import "fmt"

func main() {
	//循环开始前，会执行且仅会执行一次初始化语句 i := 0;；这比在循环之前声明更为简短
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
}
