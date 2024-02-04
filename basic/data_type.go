package main

import (
	"fmt"
	"unsafe"
)

var (
	commonVal  int
	commonTag  string
	commonBool bool = true
)

func main() {
	var b bool = true
	var a bool = false
	var c bool
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(c)

	var s string = "text demo!"
	var e string
	fmt.Println(s)
	fmt.Println(e)

	var n, m int = 1, 2
	var k int
	fmt.Println(n, m, n+m)
	fmt.Println(k)

	var p *int
	var arr []int
	var mp map[string]int
	var ch chan int
	var f func(string) int
	var err error
	fmt.Println(p, arr, mp, ch, f, err)

	// declare operation :=
	val := 1
	str := "string"
	fmt.Println(val)
	fmt.Println(str)

	fmt.Println(commonVal, commonTag, commonBool)

	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("The area is  : %d\n", area)

	// Print constant value.
	const (
		constA = "abc"
		constB = len(constA)
		constC = unsafe.Sizeof(constA)
	)
	println(constA, constB, constC)

	const (
		iotaA = iota
		iotaB = iota
		iotaC
		iotaD = "new word"
		iotaE = 838
		iotaF
		iotaG
	)
	fmt.Println(iotaA, iotaB, iotaC, iotaD, iotaE, iotaF, iotaG)
}
