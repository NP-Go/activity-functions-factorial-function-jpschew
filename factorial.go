package main

import "fmt"

//Declare factorial function here
func factorial(num int) int {
	if num <= 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

func main() {
	//Insert Code here
	var userInput int
	fmt.Println("Please enter a number: ")
	fmt.Scan(&userInput)
	result := factorial(userInput)
	fmt.Println(result)
}
