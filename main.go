package main

import "fmt"

func main() {
	var (
		firstNum int
		//secondNum int
		//operation string
	)
	fmt.Println("Input a number: ")
	n, err := fmt.Scan(&firstNum)
	if err != nil {
		fmt.Printf("it is not a number: %v", n)
		return
	}
	fmt.Println(firstNum)
}
