package main

import "fmt"

func main() {

	var x int = 10
	var y int = 20

	var res int = add(x, y)
	var res2 int = sub(x, y)
	var res3 int = multi(x, y)
	//var res4 int = div(x, y)

	fmt.Println(res)
	fmt.Println(res2)
	fmt.Println(res3)
	//fmt.Println(res4)
}

func add(a int, b int) int {

	var c int
	c = a + b
	return c

}

func sub(a int, b int) int {

	var c int
	c = a - b
	return c

}
func multi(a int, b int) int {

	var c int
	c = a * b
	return c

}

// func div(a, b int) int {

// 	var c int
// 	c = a / b
// 	return c

// }
