package main

func main() {
	i:=0
HERE:
	println(i)
	i++
	if i==5 {
		return
	}
	goto HERE
}
