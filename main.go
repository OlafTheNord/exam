package main

import (
	"exam_basic/calc"
	"fmt"
	"strings"
)

func main() {
	var (
		firstNum  int
		secondNum int
		operation string
		result    float64
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

	for strings.ToLower(strings.Trim(operation, " ")) != "stop" {
		fmt.Println("Input type of operation (+, -, *, /, %, avg, history, undo) or \"stop\" to exit: ")
		fmt.Scan(&operation)
		switch operation {
		case "+":
			result = float64(calc.Sum(firstNum, secondNum))
			fmt.Printf("Sum is: %v\n", calc.Sum(firstNum, secondNum))
		case "-":
			result = float64(calc.Subtract(firstNum, secondNum))
			fmt.Printf("Subtract is: %v\n", calc.Subtract(firstNum, secondNum))
		case "*":
			result = float64(calc.Multiply(firstNum, secondNum))
			fmt.Printf("Multiple is: %v\n", calc.Multiply(firstNum, secondNum))
		case "/":
			result = calc.Divide(firstNum, secondNum)
			fmt.Printf("Divide is: %v\n", result)
		case "%":
			result = calc.DivideWithReminder(firstNum, secondNum)
			fmt.Printf("Remind from divide is: %v\n", result)
		case "undo": //TO DO
			fmt.Println("To do...")
		case "avg": //TO DO
			fmt.Println("To do...")
		case "history": //TO DO
			fmt.Println("To do...")
		default:
			continue
		}
	}
}
