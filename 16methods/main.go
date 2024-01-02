package main

import (
	"fmt"
)

func main() {
	var Mihir User
	Mihir.name = "mihir"
	Mihir.age = 21
	Mihir.isActive = true
	Mihir.getStatus()

	// bufio.NewReader(os.Stdin);

}

type User struct {
	name     string
	age      int
	isActive bool
}

func (u User) getStatus() {
	fmt.Println(u.isActive)
}
