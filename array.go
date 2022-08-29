package main

import "fmt"

func main() {
	arrays()
}

// arrays are generally not preferred within go
// // slices are more poplular to use in go

// // variables with array, length with array it has two types

func arrays() {

	// 	var a, b, c int

	// 	fmt.Println(a, b, c)

	// 	// variables with array
	// 	var arr [5]int
	// 	fmt.Println(arr)
	// 	fmt.Println(len(arr))

	// 	//	length with array it has two types

	// 	var arr2 [5]int
	// 	arr2 = arr
	// 	arr2[4] = 70
	// 	fmt.Println(len(arr), arr)

	// 	//printing arrays - array for loop in go // for loop there is no while loop

	// 	for i := 0; i < len(arr2); i++ {
	// 		fmt.Println("Array	Index =", i, " ", arr2[i])
	// 	}

	//arrays initialization using composite literls it is very important *

	//arr := [3]int{1, 2, 3}
	arr := [...]int{1, 2, 2, 4, 5, 7, 6, 8, 9}

	fmt.Println(arr)

}
