package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	slice := make([]int, 3)

	scanner := bufio.NewScanner(os.Stdin)
	i := 0

	for {
		fmt.Print("Enter an integer (X to exit): ")
		scanner.Scan()
		input := scanner.Text()

		if input == "X" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: Invalid input!")
			break
		}

		added := false
		for j := 0; j < len(slice); j++ {
			if slice[j] >= num {
				if i < 3 {
					slice = append(slice[1:j], append([]int{num}, slice[j:]...)...)
				} else {
					slice = append(slice[:j], append([]int{num}, slice[j:]...)...)
				}
				added = true
				break
			}
		}

		if !added {
			if i < 3 {
				slice = append(slice[1:], num)
			} else {
				slice = append(slice, num)
			}
		}

		i++
		fmt.Println(slice)
	}
}
