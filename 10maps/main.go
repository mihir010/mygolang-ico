package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["cpp"] = "cplusplus"

	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("%v --> %v\n", key, value)
	}
}
