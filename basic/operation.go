package main

import "fmt"

func main() {
	calculation()
	relation()
	logic()
	bitOp()
	assignOp()
	addrOp()
}

func calculation() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Println(c)
	c = a - b
	fmt.Println(c)
	c = a * b
	fmt.Println(c)
	c = a / b
	fmt.Println(c)
	c = a % b
	fmt.Println(c)
	a++
	fmt.Println(a)
	b--
	fmt.Println(b)

	fmt.Println()
}

func relation() {
	var a int = 21
	var b int = 10
	var c int = 21

	if a == c {
		fmt.Println("a equals to c.")
	} else {
		fmt.Println("a doesn't equals to c.")

	}
	if a < b {
		fmt.Println("a is less than b.")
	} else {
		fmt.Println("a is not less than b.")
	}

	if a > b {
		fmt.Println("a is greater than b.")
	} else {
		fmt.Println("a is not greater than b.")
	}

	if a <= c {
		fmt.Println("a is less or equals to c.")
	}
	if a >= c {
		fmt.Println("a is greater or equals to c.")
	}

	fmt.Println()
}

func logic() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Println("Both are true.")
	}
	if a || b {
		fmt.Println("a is true or b is true.")
	}
	a = false
	b = true
	if a && b {
		fmt.Println("Both are true.")
	} else {
		fmt.Println("a is false or b is false.")
	}
	if !(a && b) {
		fmt.Println("Either a or b is false")
	}
}

func bitOp() {

	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Println(c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Println(c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Println(c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Println(c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Println(c)

	fmt.Println()
}

func assignOp() {
	var a int = 21
	var c int

	c = a
	fmt.Println(c)

	c += a
	fmt.Println(c)

	c -= a
	fmt.Println(c)

	c *= a
	fmt.Println(c)

	c /= a
	fmt.Println(c)

	c = 200

	c <<= 2
	fmt.Println(c)

	c >>= 2
	fmt.Println(c)

	c &= 2
	fmt.Println(c)

	c ^= 2
	fmt.Println(c)

	c |= 2
	fmt.Println(c)

	fmt.Println()
}

func addrOp() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Type of b: %T\n", b)
	fmt.Printf("Type of c: %T\n", c)

	ptr = &a
	fmt.Printf("The value of a: %d\n", a)
	fmt.Printf("The pointer value of *ptr: %d\n", *ptr)
	fmt.Printf("The pointer address of *ptr: %p\n", ptr)
	fmt.Printf("The pointer address of *ptr: %T\n", ptr)
}
