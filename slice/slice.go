package main

import "fmt"

func main() {
	// var s = make([]int, 0, 5)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))
	// s = append(s, 1)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	var k []int
	k = append(k, 1)
	fmt.Println(len(k))
	fmt.Println(cap(k))
	f := []int{}
	f = append(f, 4)
	fmt.Println(f)
	var nums = [5]int{1, 2, 3, 4, 5}
	sli := nums[0:3]
	fmt.Println(sli)
}
