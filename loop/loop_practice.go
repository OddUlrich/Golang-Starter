package main

import "fmt"

func main() {
	multiplicationChart()
}

func multiplicationChart() {
	for m := 1; m < 10; m++ {
		for n := 1; n < 10; n++ {
			fmt.Printf("%d*%d=%d\t", n, m, n*m)
		}
		fmt.Println()
	}
}
