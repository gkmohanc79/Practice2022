// package main

// import "fmt"

// func main() {
// 	var x, y = 35, 7

// 	fmt.Printf("x + y = %d\n", x+y)
// 	fmt.Printf("x - y = %d\n", x-y)
// 	fmt.Printf("x * y = %d\n", x*y)
// 	fmt.Printf("x / y = %d\n", x/y)
// 	fmt.Printf("x mod y = %d\n", x%y)

// 	x++
// 	fmt.Printf("x++ = %d\n", x)

// 	y--
// 	fmt.Printf("y-- = %d\n", y)
// }

package main

import "fmt"

func main() {
	var x, y = 15, 25
	x = y
	fmt.Println("= ", x)

	x = 15
	x += y
	fmt.Println("+=", x)

	x = 50
	x -= y
	fmt.Println("-=", x)

	x = 2
	x *= y
	fmt.Println("*=", x)

	x = 100
	x /= y
	fmt.Println("/=", x)

	x = 40
	x %= y
	fmt.Println("%=", x)
}
