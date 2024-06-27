package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var op string
	fmt.Printf("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Printf("Enter operation(+ , - , * , / ):")
	fmt.Scanln(&op)
	fmt.Printf("Enter first number:")
	fmt.Scanln(&num2)
	result := 0.0

switch op {
case "+":
	result = num1 + num2
case "-":
	result = num1 - num2
case "*":
	result = num1 * num2
case "/":
	if num2 != 0{
		result = num1 / num2
	} else{
		fmt.Printf("Can not divide by zero")
		return
	}
default:
	fmt.Printf("Invalid operator")
	return

}

fmt.Printf("Result: %.2f\n", result)


}


