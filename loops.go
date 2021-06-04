package main

import "fmt"

func main() {
	for a := 0; a < 5; a++ {
		if a == 3 {
			fmt.Println("has 3")
		}
		for b := 0; b < 5; b++ {
			fmt.Println(a, b)
		}
	}
}
