package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name    string   `json:"name"`
	Price   int      `json:"price"`
	Website string   `json:"website"`
	Tags    []string `json:"tags,omitempty"`
}

type course2 struct {
	name    string
	price   int
	website string
	tags    []string
}

func main() {
	// courseDetails := course2{"ReactJS", 299, "reactJS.com", []string{"web-dev", "frontend"}}
	// courseDetails.encodeString2()
	encodeString()
}

func encodeString() {
	courseNames := []course{
		{"ReactJS", 299, "reactJS.com", []string{"web-dev", "frontend"}},
		{"NodeJS", 399, "nodeJS.com", []string{"web-dev", "backend"}},
		{"MERN", 799, "mernstack.com", []string{"web-dev", "frontend", "backend", "full-stack"}},
		{"MEAN", 799, "meanstack.com", []string{}},
	}
	encodedFile, err := json.MarshalIndent(courseNames, "", "\t")

	if err != nil {
		panic(err)
	}

	// fmt.Printf("%s\n", encodedFile)
	fmt.Println(string(encodedFile))

}

func (course2) encodeString2() {

}
