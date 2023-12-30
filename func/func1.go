package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200
	var ret int
	ret = maxNum(a, b)

	fmt.Printf("Max value is : %d\n", ret)
}

func maxNum(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}
