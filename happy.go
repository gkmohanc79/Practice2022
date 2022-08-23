package main

import "fmt"

func main() {

	var x = 10
	var y = &x
	var z = &y
	//fmt.Println(&x, *y, **z)
	fmt.Println(x, &y, *z)
}
