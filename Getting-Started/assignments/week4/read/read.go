package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	scanner1 := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter file name: ")
	scanner1.Scan()
	fileName := scanner1.Text()

	filedata := make([]name, 0)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()

	scanner2 := bufio.NewScanner(file)
	for scanner2.Scan() {
		line := scanner2.Text()
		lineData := strings.Split(line, " ")
		if len(lineData[0]) > 20 {
			lineData[0] = lineData[0][:20]
		}
		if len(lineData[1]) > 20 {
			lineData[1] = lineData[1][:20]
		}
		filedata = append(filedata, name{lineData[0], lineData[1]})
	}

	for i, v := range filedata {
		fmt.Println(strconv.Itoa(i+1)+".", "First name:", v.fname+",", "Last name:", v.lname)
	}
}
