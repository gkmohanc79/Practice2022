package main

import "fmt"

// func main() {

// 	arrays()
// }

// func main() {
// 	goslices()
// }

// func main() {

// }

// func main() {

// 	var x int = 10
// 	var y = &x
// 	var z = y
//    // pointer of pointer to the address*
// 	fmt.Println("\n", x, y, *y, &x, &y, z, &z, *z)
// 	fmt.Println("\n Welcome to America")
// }

// Arthemetic calculations .......

// func main() {

// 	var x, y, k = 5, 10, 20
// 	Res := add(x, y, k)
// 	Res1 := sub(x, y, k)
// 	Res2 := div(x, y, k)
// 	Res3 := mul(x, y, k)
// 	Res4 := mod(x, y, k)

// 	fmt.Println("\n", Res, " \n ", Res1, "\n  ", Res2, "\n  ", Res3, " \n ", Res4)

// 	// fmt.Println("Index => Addition value is: ", " ", Res, "Address of =>  ", &Res)
// 	// fmt.Println("Index =>Substraction value is: ", " ", Res1)
// 	// fmt.Println("Division value is: ", Res2)
// 	// fmt.Println("Multiplication value is: ", Res3)
// }
// func add(x, y, k int) int {
// 	var z int
// 	z = x + y + k
// 	return z
// }
// func sub(x, y, k int) int {
// 	var z int
// 	z = x - y - k
// 	return z
// }

// func div(x, y, k int) int {
// 	var z int
// 	z = x / y / k
// 	return z
// }
// func mul(x, y, k int) int {
// 	var z int
// 	z = x * y * k
// 	return z
// }
// func mod(x, y, k int) int {
// 	var z int
// 	z = x % y % k
// 	return z
// }

const s string = "constant"

func main() {
	const n = 500000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(math.sin(n))
}
