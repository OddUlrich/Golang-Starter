package main

import "fmt"

func main() {
	var i, j, k int
	balance := [5]float32{100.0, 22.0, 331.4, 17.0, 5044.0}

	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	balance2 := [...]float32{120.0, 233.0, 322.44, 7.20, 5110.0}
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.320, 3: 721.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}
}
