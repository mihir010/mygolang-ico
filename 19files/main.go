package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "hello hello world!!!"

	file, err := os.Create("./text.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println(length)
	defer file.Close()

	readFile()
}

func readFile() {
	file, err := os.ReadFile("./text.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))
}
