package main

import (
	"fmt"
	"io/ioutil"
)

func fileAccessIO() {
	//  File Access with ioutil package with open, read, write, close, seek, and other functions

	dat, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dat))

	err = ioutil.WriteFile("test.txt", []byte("Hello World!"), 0644)
	if err != nil {
		panic(err)
	}

	dat, err = ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dat))
}
