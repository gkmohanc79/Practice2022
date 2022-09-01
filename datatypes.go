package main

import (
	"fmt"
)

func main() {
	// intergers
	var a int = -2
	var b int = 60
	c := 3 // it shows without giving Variable declaration through : we can call

	fmt.Println(" Multiplication of a and b is:", (a * b))
	fmt.Println("Random calculation:", (c + 4 - a*b))
	//float

	d := 5 + 3.2
	fmt.Println(" Float value of d is :", d)

	// complex
	e := 20 + 27i
	fmt.Println("Complex value is:", e)
	// string value with concatenate

	f := " Smiley is beautiful"
	g := " her is bueaty"
	fmt.Println(f + g)

	//Boolean

	//  Var h bool
	//  i := true
	//  fmt.Println("H: ", h)
	//  fmt.Println("I: ", i)

	// Type of conversion

	j := 7.2 // 7
	//var k float32 = j

	var k float32 = float32(j)
	fmt.Println("Value of k is: ", k)
}
