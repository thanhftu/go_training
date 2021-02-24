package main

import "fmt"

var a string

// G->O->O
func main() {
	a = "G"
	fmt.Println(a)
	f1()
}
func f1() {
	a := "O"
	fmt.Println(a)
	f2()
}

func f2() {
	// this "a" is from global
	fmt.Println(a)
}
