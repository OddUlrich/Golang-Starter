package main

import "fmt"

func main() {
	var a = 10

LOOP:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}
