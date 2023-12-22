package main

import "fmt"

func main() {
	sumNumber()
	selfSum()
	rangeLoop()
	skipRangeLoop()

	inLoop()
}

func sumNumber() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func selfSum() {
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	// looks like while expression.
	anoSum := 1
	for anoSum <= 10 {
		anoSum += anoSum
	}
	fmt.Println(anoSum)
}

func infinityLoop() {
	sum := 0
	for {
		sum++
		fmt.Println(sum)
	}
}

func rangeLoop() {
	strings := []string{"google", "baidu", "yahoo"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	nums := [6]int{1, 2, 3, 5}
	for i, x := range nums {
		fmt.Println(i, x)
	}
}

func skipRangeLoop() {
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	map1[5] = 5.0
	map1[6] = 6.0

	fmt.Println("key - value")
	for key, val := range map1 {
		fmt.Println(key, val)
	}

	fmt.Println("key")
	for key := range map1 {
		fmt.Println(key)
	}

	fmt.Println("value")
	for _, val := range map1 {
		fmt.Println(val)
	}
}

func inLoop() {
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > i/j {
			fmt.Printf("%d is norm.\n", i)
		}
	}
}
