package main

import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	if a >= 13 && a < 19 {
		fmt.Println("Youre a teenager")
	} else {
		panic("Youre not allowed ")
	}
}
