package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("https://lco.dev/")

	if err != nil {
		panic(err)
	}

	databytes, err := io.ReadAll(response.Body)

	defer response.Body.Close()

	fmt.Printf("%T", databytes)

	// fmt.Println(string(databytes))

}
