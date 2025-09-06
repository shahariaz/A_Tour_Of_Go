package main

import (
	"fmt"
	"sync"
)

//Go Routine

// func main() {
// 	for i := 0; i <= 10; i++ {
// 		go func() {
// 			fmt.Println("Counting : ", i)
// 		}()
// 	}
// 	time.Sleep(time.Second * 1)
// }

// WaitGroup

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		func(i int) {
			wg.Add(1)
			fmt.Println("Counitng : ", i)
			wg.Done()
		}(i)
		wg.Wait()
	}
}
