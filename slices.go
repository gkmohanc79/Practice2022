package main

import "fmt"

func main() {

	goslices()

}

func goslices() {
	// 	// slices are similar to arrays, but preferred
	// 	// we dont specify any size in slices

	// 	//var slice []int
	// 	//slice = []int{1, 2}

	// 	//slice = []int{1, 2, 3}

	// 	// difference between array and slices - length is not part of the type
	// 	// array := [5]int{1, 2, 3, 4, 5}

	// 	// var array1 [5]int

	// 	// array1 = array
	// 	// fmt.Println(" new array=> ", array1)

	// 	//fmt.Println(slice)
	// 	// return

	// 	// lets try with slice, for slice length is not part of type*

	// 	slice1 := []int{1, 2, 3, 4, 5, 6, 6, 4, 7}

	// 	var slice2 []int
	// 	slice2 = slice1
	// 	fmt.Println("New slice =>", slice2)

	// 	//interacting through slice - range

	// 	slice3 := []int{1, 2, 3, 4, 5}

	// 	for i, v := range slice3 {

	// 		fmt.Println(" Index => ", i, " value => ", v)
	// 		//return //-    it shows only one index values

	// 		// in Go if you declare the variable u must to be use it

	// 		slice4 := []int{1, 2, 3, 4, 5}

	// 		fmt.Println(" length of slice4: ", len(slice4))
	// 		fmt.Println(" capacity of slice4: ", cap(slice4))
	// 		// in slices, you can append additional elements
	// 		// length and capacity of slice ( array also will have length and capacity)
	// 		slice4 = append(slice4, 9)

	// 		fmt.Println(" length of slice4: ", len(slice4))
	// 		fmt.Println(" capacity of slice4: ", cap(slice4))

	// 		// in slices, you can append multiple elements

	// 		slice4 = append(slice4, 1, 2, 3, 4)

	// 		fmt.Println(" length of slice4: ", len(slice4))
	// 		fmt.Println(" capacity of slice4: ", cap(slice4))

	// 		return
	// 	}

	// let's talk about Make function

	slice5 := make([]int, 5)

	fmt.Println(slice5)

	fmt.Println(" length of slice5: ", len(slice5))s
	fmt.Println(" capacity of slice5: ", cap(slice5))
	return
}
