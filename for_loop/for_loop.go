package main

import "fmt"

func main() {
	numbers := [4]int{1, 3, 45, 6}
	for index, value := range numbers {
		if value == 45 {
			continue
		}
		if value == 3 {
			break
		}
		fmt.Printf("Index: %d, value: %d\n", index, value)
	}
}
