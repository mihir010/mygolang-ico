package main

import (
	"fmt"
	"net/url"
)

var myurl string = "https://pkg.go.dev/net/url"

func main() {
	result, _ := url.Parse(myurl)

	fmt.Println(result.Host)
	// fmt.Println(result.Query())
	qs := result.Query()

	for _, val := range qs {
		fmt.Println(val)
	}

	// partsOfURL := &url.URL{
	// 	Scheme: ,
	// }

}
