package main

import "fmt"

func main() {

	var x = 10
	var y = &x

	fmt.Println(&y, *y, x, &x, y)
}
