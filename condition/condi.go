package condition

import "fmt"

func condition() {
	a := 23
	if a >= 18 {
		fmt.Println("He can get a driving licese now")
	} else if a == 17 {
		fmt.Println("He is 17")
	} else {
		fmt.Println("He can't get license")
	}

}
