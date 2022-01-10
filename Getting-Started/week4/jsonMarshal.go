package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Addr  string
	Phone string
}

func jsonMarshal() {
	p1 := &Person{Name: "John", Addr: "USA", Phone: "123"}

	barr, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(barr))

	err = json.Unmarshal(barr, &p1)
	if err != nil {
		panic(err)
	}

	fmt.Println(*p1)
}
