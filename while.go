package main

import "fmt"

func main() {
	sum := 1
	for sum < 200 {

		fmt.Println(sum)
		sum += sum
	}
	array := []int{1, 2, 3, 4, 5, 6, 7}
	mapping := map[int]string{
		1: "sujai",
		2: "vijay",
		3: "ganesh",
		4: "dhinesh",
	}

	for i, k := range mapping {
		fmt.Println(i, "", k)
	}

		sum := 0
		var total int
		for _, k := range array {
			sum += k
			total = sum

		}
		fmt.Println(total)
	}

		array := []int{1, 2, 3, 4, 5, 6, 7}
		mapping := map[int]string{
			1:"sujai",
			2: "vijay",
			3: "ganesh",
			4: "dhinesh",
		}

		for i, k := range array {
			fmt.Println(i, "", k)

	Loop:
		fmt.Println("youre not eligible")
		fmt.Println("enter youre age")
		var age int
		fmt.Scanln(&age)
		if age < 18 {
			goto Loop
			// fmt.Println("Youre not eligible")
		} else {
			fmt.Println("youre eligible")
		}

	Controlling the flow of Loops

	for a := 0; a < 6; a++ {
		// fmt.Println(a)
		// if a == 2 {
		// 	break
		// }
		if a == 4 {
			continue
		}
		fmt.Println(a)
	}

}
