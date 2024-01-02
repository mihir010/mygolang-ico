package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	temp := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println(temp.Intn(10))
	fmt.Println(temp)
	// fmt.Printf("%T", temp.Int())
}
