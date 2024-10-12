package main

import "fmt"

// 2. Basic Calculator
// Write a program that takes two numbers as input from the user and performs addition, subtraction, multiplication, and division operations on them. Print the results to the console.
// ,

func main() {

	var first, second float64
	fmt.Print("Enter first number: ")
	_, err := fmt.Scan(&first)
	if err != nil {
		fmt.Println("Please enter a valid number")
		return
	}
	fmt.Print("Enter second number: ")
	_, err = fmt.Scan(&second)
	if err != nil {
		fmt.Println("Please enter a valid number")
		return
	}
	fmt.Print("Choose Operation between + - * / : ")
	var operation string
	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Please enter a valid operation")
		return
	}

	if operation != "+" && operation != "-" && operation != "*" && operation != "/" {
		fmt.Println("Please enter a valid operation")
		_, err = fmt.Scan(&operation)
		if err != nil {
			fmt.Println("Please enter a valid operation")
			return
		}
	}

	switch operation {
	case "+":
		fmt.Println(first + second)
	case "-":
		fmt.Println(first - second)
	case "*":
		fmt.Println(first * second)
	case "/":
		fmt.Println(first / second)
	}
}
