package num

import (
	"fmt"
	"math/rand"
)

// Exported function (can be used from other packages)
func Println() {
	fmt.Println("My favorite number is", rand.Intn(10))

}
