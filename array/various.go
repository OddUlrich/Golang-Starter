package main

import "fmt"

func main() {
	var animals [][]string

	row1 := []string{"fish", "shark", "eel"}
	row2 := []string{"bird"}
	row3 := []string{"lizard", "salamander"}

	animals = append(animals, row1)
	animals = append(animals, row2)
	animals = append(animals, row3)

	for i := range animals {
		fmt.Printf("Row: %v\n", i)
		fmt.Println(animals[i])
	}
}
