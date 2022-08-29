package main

import "fmt"

func main() {

	var x = 10
	var y = &x
	//var z = &y
	//fmt.Println(&x, *y, **z)
	fmt.Println(y)
}

// func main() {

// 	var i = 55
// 	var j = 66.19

// 	add := float(i) + j

// 	fmt.Println("add")
// }

// func add(i int, j int) int {

// 	var c int
// 	c = i + j
// 	return c
// }
