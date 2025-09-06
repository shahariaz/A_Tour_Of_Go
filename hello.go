package main

import (
	num "a_go/package"
	"a_go/variable"
	"fmt"
)

var (
	x = "hello"
)

func main() {
	fmt.Println("Hello, 世界")
	num.Println()
	num.Find()
	variable.Variable()
	test()

}
