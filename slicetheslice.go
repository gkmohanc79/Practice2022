package main

import "fmt"

func main() {

	slicetheslices()
}

func slicetheslices() {

	slice := make([]int, 6)

	// for i := 0; i < len(slice); i++ {
	// 	slice[i] = i + 10

	// }

	// fmt.Println(slice)

	for i := 0; i < len(slice); i++ {
		slice[i] = i + 10
		//	s2 := slice
		//s2 := slice[2:5] // index values include and exlude the value
		s2 := slice[2:]

		fmt.Println(s2)
	}
}
