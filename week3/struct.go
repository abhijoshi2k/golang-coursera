package main

import "fmt"

type Person struct {
	name string
	age  int
}

func structs() {
	var p1 Person
	p1.name = "John"
	p1.age = 30
	fmt.Println("p1:", p1)

	p2 := Person{name: "Mary", age: 25}
	fmt.Println("p2: name:", p2.name, "age:", p2.age)

	p3 := Person{name: "Mike"}
	fmt.Println("p3:", p3)

	p4 := Person{}
	fmt.Println("p4:", p4)

	// Other way to initialize struct
	p5 := new(Person)
	p5.name = "John"
	p5.age = 30
	fmt.Println("p5:", p5)
}
