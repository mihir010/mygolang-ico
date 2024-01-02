package main

import "fmt"

func main() {
	var list [4]string
	list[0] = "mihir"
	list[1] = "aaryan"
	list[2] = "naincy"

	fmt.Println(list)
	fmt.Println(len(list))

	var sub = [3]string{"sp", "dat", "ai"}

	fmt.Println(sub)
}
