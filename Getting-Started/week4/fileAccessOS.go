package main

import (
	"fmt"
	"os"
)

func fileAccessOS() {
	// Rename the file
	error := os.Rename("test.txt", "test2.txt")
	if error != nil {
		panic(error)
	}

	// Create a new file and write to it
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("Hello World!")

	// Open the file we created before
	f, err = os.OpenFile("test.txt", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file
	b := make([]byte, 100)
	n, err := f.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b[:n]))

	// Seek ahead by 5 bytes
	f.Seek(5, 0)

	// Read the file
	b = make([]byte, 100)
	n, err = f.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b[:n]))
}
