package main

import "fmt"

func array() {
	var a [5]int
	var b [5]int = [5]int{1, 2, 3, 4, 5}
	c := [...]string{"a", "b", "c", "d", "e"}
	d := [4]string{"geek", "gfg", "Geeks1231", "GeeksforGeeks"}
	a[4] = 100
	fmt.Println("element at index 1:", a[1])
	fmt.Println("element at index 4:", a[4])
	fmt.Println("length of array:", len(a))
	fmt.Println("array b:", b)
	fmt.Println("array c:", c)
	fmt.Println("array d:", d)

	fmt.Println("----------------------------------------------------")

	for i, v := range c {
		fmt.Println("index:", i, "value:", v)
	}
}
