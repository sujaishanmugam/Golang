package main

import "fmt"

type phone struct {
	ram   int
	color string
	price float64
}
type company struct {
	phone
	name string
}

func main() {
	// ptr := new(int)
	// *ptr = 90
	// fmt.Println(ptr)
	y := company{phone: phone{ram: 90, color: "red", price: 5000.65}, name: "iphone"}
	x := phone{ram: 90, color: "red", price: 5000.65}
	fmt.Printf("%v,%v", x, y)

}
