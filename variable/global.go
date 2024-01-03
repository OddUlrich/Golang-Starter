package main

import "fmt"

var g int

func main() {

	var a, b int

	/* 初始化参数 */
	a = 10
	b = 20
	g = a + b

	fmt.Printf("结果： a = %d, b = %d and g = %d\n", a, b, g)
}
