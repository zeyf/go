package main

import (
	"fmt"
)

func print_list(list []int, reverse bool) {
	for _, value := range list {
		if reverse {
			defer fmt.Printf("%d ", value)
		} else {
			fmt.Printf("%d ", value)
		}
	}
}

func main() {
	var list []int = []int{1, 2, 3, 4, 5}

	fmt.Println("Printing in normal order...")
	print_list(list, false)
	fmt.Println()

	fmt.Println("Printing in reverse order...")
	print_list(list, true)
	fmt.Println()
}
