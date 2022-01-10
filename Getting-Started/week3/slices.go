package main

import "fmt"

func slice() {
	// slice is extendable
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"} // array
	s1 := arr[1:3]
	s2 := []string{"h", "i", "j"} // slice - covers whole array
	s3 := make([]int, 3)          // slice - covers whole array
	s4 := make([]int, 3, 5)       // slice - covers starting from index 0 and has length 3

	s1[0] = "x"
	s2[0] = "y"

	// Changing length of slice beyond its capacity has time penalty!!!!!

	fmt.Println("arr:", arr)
	// length and capacity of slice
	s1[1] = "z"
	s1 = append(s1, "k", "l", "m", "n", "o") // slice will get dereferrenced and then appended
	fmt.Println("len(s1):", len(s1))
	fmt.Println("cap(s1):", cap(s1))
	fmt.Println("arr:", arr)

	fmt.Println("slice s2:", s2)

	fmt.Println("slice s3:", s3)

	s4 = append(s4, 1, 2, 3, 4) //
	fmt.Println("slice s4:", s4)
}
