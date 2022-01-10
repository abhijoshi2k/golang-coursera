package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	person := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter name: ")
	scanner.Scan()
	person["name"] = scanner.Text()

	fmt.Print("Enter address: ")
	scanner.Scan()
	person["address"] = scanner.Text()

	data, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(data))
}
