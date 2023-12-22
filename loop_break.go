package main

import (
	"fmt"
	"time"
)

func main() {
	simpleBreak()
	breakLabel()
	switchBreak()
	selectBreak()
	returnSelect()
}

func simpleBreak() {
	var a int = 10

	for a < 20 {
		fmt.Println(a)
		a++
		if a > 15 {
			break
		}
	}
}

func breakLabel() {
	fmt.Println("---------  break label  ----------")

label:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Print(i, j, "  ")

			if j == 3 {
				break label
			}
		}
	}

}

func switchBreak() {
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Printf(day)
		break
	case "Sunday":
		fmt.Printf("Sunday")
	}
}

func selectBreak() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 2
	}()

	select {
	case <-ch1:
		fmt.Println("Received from ch1.")
	case <-ch2:
		fmt.Println("Received from ch2.")
		// call break, then wait all channel completing tasks.
		break
	}
}

func returnSelect() {
	ch := make(chan int)

	go process(ch)

	time.Sleep(2 * time.Second)
	ch <- 1
	time.Sleep(1 * time.Second)
	ch <- 3
	time.Sleep(1 * time.Second)
	ch <- 5
	time.Sleep(1 * time.Second)
	ch <- 7

	time.Sleep(2 * time.Second)
}

func process(ch chan int) {

	for {
		select {
		case val := <-ch:
			fmt.Println("Received value:", val)
			if val == 5 {
				// Stop select immediately
				return
			}
		default:
			fmt.Println("No value received yet.")
			time.Sleep(500 * time.Millisecond)
		}
	}

}
