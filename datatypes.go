package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// intergers
	var a = -2
	var b = 60
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

	// zero value of Golang data types
	// int  is 0, float is 0.0, string is "", boolean is "false"

	f := " Smiley is beautiful"
	g := " she  is beauty"
	fmt.Println(f + g)
	fmt.Println("The time is ", time.Now())
	//Boolean

	//  Var h bool
	//  i := true
	//   fmt.Println("H: ", h)
	//   fmt.Println("I: ", i)

	// Type of conversion

	j := 7.2 // 7
	//var k float32 = j

	var k float32 = float32(j)
	fmt.Println("Value of k is: ", k)
	// for loop

	for l := 5; l <= 10; l++ {
		fmt.Println(l)
	}
	//decision making if & else or switch

	// var age = 19
	// if age != 18 {

	// 	fmt.Println(" you can win")
	// } else {
	// 	fmt.Println(" you can go India")
	// }

	// var name = krishna
	// if name != krishna {

	// 	fmt.Println(" you can win")
	// } else {
	// 	fmt.Println(" you can go India")
	// }

	age := 28

	switch age {

	case 16:
		fmt.Println("Prepare for collage")
	case 18:
		fmt.Println("Prepare for girl")
	case 20:
		fmt.Println("Prepare for job")
	default:
		fmt.Println("Prepare for alone")
	}
	fmt.Println("Now you have %g problems", math.Sqrt(57))

}
