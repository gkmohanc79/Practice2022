// package main

// import "fmt"

// func swap(x, y, z string) (string, string, string) {
// 	return x, y, z
// }

// func main() {
// 	a, b, c := swap("hello", "world", " ")
// 	fmt.Println(a, b, c)
// }

// package main

// import (
// 	"fmt"
// 	"math/cmplx"
// )

// var (
// 	ToBe   bool       = false
// 	MaxInt uint64     = 1<<64 - 1
// 	z      complex128 = cmplx.Sqrt(-5 + 12i)
// )

// func main() {
// 	fmt.Printf("Type: %T Value: %v\n", ToBe)
// 	fmt.Printf("Type: %T Value: %v\n", MaxInt)
// 	fmt.Printf("Type: %T Value: %v\n", z)
// }

// Golang program to illustrate
// the strings.Contains() Function
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {

// 	// using the function
// 	fmt.Println(strings.Contains("GeeksforGeeks", "for"))
// 	fmt.Println(strings.Contains("A computer science portal", "science"))

// }

// Type conversions
// Unlike in C, in Go assignment between items of different type requires an explicit conversion.
// Try removing the float64 or uint conversions in the example and see what happens.
package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, z)
}
