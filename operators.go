package main

import "fmt"

func main() {
	a := 15
	b := 15
	fmt.Printf("%v", a+b)
	a += b
	fmt.Printf("%v", a)
}
