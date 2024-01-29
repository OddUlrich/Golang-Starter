package main

import "fmt"

func main() {
	var values [][]int

	// Use append function
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	// Print the whole row
	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	fmt.Println("第一个元素为：")
	fmt.Println(values[0][0])

	// Initialize bi-dimensional matrix
	sites := [2][2]string{}

	sites[0][0] = "JetBrains"
	sites[0][1] = "Oracle"
	sites[1][0] = "DIDI"
	sites[1][1] = "X"
	fmt.Println(sites)

	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
