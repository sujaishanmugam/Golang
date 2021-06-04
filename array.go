package main

import "fmt"

func main() {
	// 1D array
	var a = []int{0, 1, 2, 3, 4, 6}
	var b = a[0:]
	fmt.Println(b)

	// 2D array
	var a = []int{0, 1, 2, 3, 4, 6,}
	var b = [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(b)

	// struct in arrays
	b := []struct {
		name string
		i    bool
	}{
		{"sujai", true},
		{"vijay", false},
// 	}
	fmt.Println(b)
}
