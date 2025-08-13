package main

import "fmt"

func main() {
	var (
		firstNum  int
		secondNum int
		operation string
	)
	fmt.Println("Input a first number: ")
	_, err := fmt.Scan(&firstNum)
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	fmt.Println("Input a second number: ")
	_, err = fmt.Scan(&secondNum)
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	fmt.Println("Input type of operation (+, -, *, /, avg, history, undo): ")
	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	println(firstNum, operation, secondNum)
}
