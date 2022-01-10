package main

import "fmt"

func main() {
	fmt.Println("Enter Floating Point Number")
	var floatNum float64
	_, err := fmt.Scan(&floatNum)
	if err != nil {
		fmt.Println("Error: Invalid Input")
		return
	}
	fmt.Println("Entered Number is:", floatNum)
	fmt.Printf("Truncated Number is: %d", int(floatNum))
}
