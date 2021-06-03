// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// test funtions

// func main() {
// 	// var score int = 4
// 	fmt.Println("hello")
// }

// variables

// func main() {
// 	var tonyStark string = "I am tonyStark"
// 	fmt.Println(tonyStark)

// 	var dom string
// 	dom = "I am Dom"
// 	fmt.Println(dom)

// 	Paul := "I am Paul"
// 	fmt.Println(Paul)

// 	tonyStarkRating := 10.0
// 	fmt.Printf("%v, %T", tonyStarkRating, tonyStarkRating)

// 	var Ironman, CaptAmerica string = "I am Ironman", "I am Capt America"
// 	fmt.Println(Ironman, CaptAmerica)

// 	var (
// 		spideman = "I am Spiderman"
// 		age      = 18
// 		powers   = "web slings, spidy sense"
// 	)
// 	fmt.Println(spideman, age, powers)
// }

// takeInput

// func main() {
// 	var myString string
// 	fmt.Scanln(&myString)
// 	fmt.Println(myString)

// 	var name string = "sujai"
// 	var a, b = fmt.Println(name)
// 	fmt.Println(a, b)

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter youre full name: ")
// 	myName, _ := reader.ReadString('\n')
// 	fmt.Println(myName)

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter your rating: ")
// 	myRating, _ := reader.ReadString('\n')
// 	mynumRating, _ := strconv.ParseInt(strings.TrimSpace(myRating), 64)
// 	fmt.Println(mynumRating + 2)
// }
