package main

import "fmt"

func main() {

	var choice int
	fmt.Println("1. Addition ")
	fmt.Println("2. Subtraction ")
	fmt.Println("3. Division ")
	fmt.Println("4. Multiplication ")
	fmt.Scan(&choice)

	var num1 int
	var num2 int

	switch choice {
	case 1:
		fmt.Println("Enter Number 1 ")
		fmt.Scan(&num1)
		fmt.Println("Enter Number 2")
		fmt.Scan(&num2)
		var add = num1 + num2
		fmt.Println("Addition of two numbers :  ", add)
	case 2:
		fmt.Println("Enter Number 1 ")
		fmt.Scan(&num1)
		fmt.Println("Enter Number 2")
		fmt.Scan(&num2)
		var add = num1 + num2
		fmt.Println("Addition of two numbers :  ", add)
	case 3:
		fmt.Println("Enter Number 1 ")
		fmt.Scan(&num1)
		fmt.Println("Enter Number 2")
		fmt.Scan(&num2)
		var add = num1 + num2
		fmt.Println("Addition of two numbers :  ", add)
	case 4:
		fmt.Println("Enter Number 1 ")
		fmt.Scan(&num1)
		fmt.Println("Enter Number 2")
		fmt.Scan(&num2)
		var add = num1 + num2
		fmt.Println("Addition of two numbers :  ", add)
	default:
		fmt.Println("Error")

	}

}
