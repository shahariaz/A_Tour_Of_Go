package main

import "fmt"

func main() {

	mp := make(map[string]int)
	mp["hello"] = 0
	mup := map[string]string{"kui": "lui"}
	mup["ui"] = "hui"

	k, ok := mup["ui"]
	fmt.Println(k)
	if !ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

}
