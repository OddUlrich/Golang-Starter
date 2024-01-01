package main

import "fmt"

func main() {
	demo1()
	demo2()
}

func demo1() {
	nextNumber := getSequence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func demo2() {
	add := func(a, b int) int {
		return a + b
	}

	result := add(3, 5)
	fmt.Println("3 + 5 = ", result)

	multiply := func(x, y int) int {
		return x * y
	}

	product := multiply(4, 6)
	fmt.Println("4 & 6 = ", product)

	calculate := func(operation func(int, int) int, x, y int) int {
		return operation(x, y)
	}

	sum := calculate(add, 2, 8)
	fmt.Println("2 + 8 = ", sum)

	difference := calculate(func(a, b int) int {
		return a - b
	}, 10, 4)
	fmt.Println("10 - 4 = ", difference)

}
