package main

import "fmt"

func getSum(a int, b int) (int, int) {
	return (a + b), (b + 1)

}

func outer() func() {
	money := 100
	age := 10
	show := func() {
		money = money * age
		fmt.Println(money)
	}
	return show
}

func function() {
	p, q := getSum(3, 5)

	hello := func() {
		fmt.Println("hello", p, q)
	}

	hello()
	ok := func() {}
	ok()

}
func main() {
	a, b := getSum(4, 5)
	outer()
	function()
	fmt.Println(a, b)

}
