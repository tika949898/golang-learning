package main

import (
	"fmt"
)

func main() {
	var student1 string = "Tikaram" //this is a string
	var student2 = "Saajan"         //This is inferred
	x := 2                          //type is inferred
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}
