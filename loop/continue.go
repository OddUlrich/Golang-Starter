package main

import "fmt"

func main() {
	basicContinue()
	labelContinue()
}

func basicContinue() {
	var a int = 10

	for a < 20 {
		if a == 15 {
			/* 跳过此次循环 */
			a = a + 1
			continue
		}
		fmt.Printf("a 的值为 : %d\n", a)
		a++
	}
}

func labelContinue() {
	fmt.Println("---- continue label ----")

label:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue label
		}
	}
}
