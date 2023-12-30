package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	fmt.Printf("Before swap a : %d\n", a)
	fmt.Printf("Before swap b : %d\n", b)

	swapReference(&a, &b)
	fmt.Printf("After swap a : %d\n", a)
	fmt.Printf("After swap b : %d\n", b)
}

func swapReference(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
