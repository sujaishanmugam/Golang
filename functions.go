package main

import "fmt"

type person struct {
	fname string
	lname string
}

func (p person) fullName() {
	fmt.Println(p.fname, "", p.lname)
}

func calc(a, b int) {
	// var a, b int
	// a = 80
	// b = 70
	c := a + b
	fmt.Println(c)
}

func newCalc(a int, b int) int {
	return a + b
}

func arguments(args ...int) (int, int) {
	a := 0
	b := 0
	for _, j := range args {
		a += j
		b++
	}
	return a, b
}

func main() {
	a := person{"John", "vijay"}
	a.fullName()
	calc(12, 2)
	fmt.Println(newCalc(2, 5))
	fmt.Println(arguments(10, 52, 55, 3, 5, 65, 6))

	func() {
		fmt.Println("i am sujai")
	}()

	s := make(chan string)

	go func() {
		s <- "i am sujai"
	}()
	text := <-s

	fmt.Println(text)
}
