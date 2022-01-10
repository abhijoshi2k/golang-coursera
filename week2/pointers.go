package main

import "fmt"

func pointers() {
	i := 11111       // i assigned to 11111
	var ip *int = &i // ip assigned to address of i
	j := *ip         // j assigned to value at address pointed to by ip
	fmt.Println(j)
}
