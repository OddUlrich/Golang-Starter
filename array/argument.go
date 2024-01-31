package main

import "fmt"

func modifyArray(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
}

func modifyArrayWithPointer(arr *[5]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = (*arr)[i] * 2
	}
}

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Original Array:", myArray)

	modifyArray(myArray)
	fmt.Println("Array after modifyArray:", myArray)

	modifyArrayWithPointer(&myArray)
	fmt.Println("Array after modifyArrayWithPointer:", myArray)
}
