package main

import "fmt"

var a = "G"

func main() {
	n() // G
	m() // O
	n() // O
}

func n() {
	fmt.Println(a)
}
func m() {
	a = "O"
	fmt.Println(a)
}
