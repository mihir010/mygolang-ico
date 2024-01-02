package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"apple", "orange", "tomato", "guava", "grapes"}

	fmt.Println(fruits)

	fruits = append(fruits, "strawberries")

	index := 2

	fruits = append(fruits[:index], fruits[index+1:]...)

	fmt.Println(fruits)

	fmt.Println(fruits[1:])

	highScores := make([]int, 4)

	highScores[0] = 239
	highScores[1] = 265
	highScores[2] = 931
	highScores[3] = 394

	highScores = append(highScores, 123, 454, 133)

	sort.Ints(highScores)
}
