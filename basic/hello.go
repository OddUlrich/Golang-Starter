package main

import "fmt"

/*
Block-comment.
*/
func main() {
	// Line-comment.
	fmt.Println("Hello, World!")
	fmt.Println("Good" + " " + "morning" + "!")

	var format = "%s %d %s."

	var date1 = "2021-10-12"
	var index1 = 1
	var action1 = "Say hello"

	var s1 = fmt.Sprintf(format, date1, index1, action1)
	fmt.Println(s1)

	fmt.Printf(format, "2022-02-21", 2, "Print result")

}
