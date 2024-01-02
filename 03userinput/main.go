package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your age: ")
	res := bufio.NewReader(os.Stdin)

	input, _ := res.ReadString('\n')

	fmt.Println("Youe age is: ", input)
	fmt.Printf("Type %T", input)
}
