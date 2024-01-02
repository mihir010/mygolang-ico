package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("whats your age")
	res := bufio.NewReader(os.Stdin)

	input, _ := res.ReadString('\n')

	numAge, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1: ", numAge+1)
	}
}
