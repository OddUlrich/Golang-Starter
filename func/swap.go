package main

import "fmt"

func main() {
	a, b := swap("Goole", "Baidu")
	fmt.Println(a, b)

	var c int = 100
	var d int = 200

	fmt.Printf("Before swap c : %d\n", c)
	fmt.Printf("Before swap d : %d\n", d)

	swapTemp(c, d)
	fmt.Printf("After swap c : %d\n", c)
	fmt.Printf("After swap d : %d\n", d)
}

func swap(x, y string) (string, string) {
	return y, x
}

func swapTemp(x, y int) int {
	var temp int

	temp = x
	x = y
	y = temp

	return temp

}
