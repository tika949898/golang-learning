// This example shows declaring variables outside of a function, with the var keyword:
package main

import (
	"fmt"
)

var a int
var b int = 2
var c = 3

func main() {
	a = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
