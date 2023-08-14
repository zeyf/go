package main

import (
	"fmt"
)

func is_even(number int) (bool, error) {
	if number%2 == 0 {
		return true, nil
	} else {
		return false, fmt.Errorf("%d is not even, it is odd.", number)
	}
}

func main() {
	var numbers []int = []int{4, 7, 10, 9, 1}

	for _, value := range numbers {
		// NOTE: If variable declarations and initializations have context for
		// downstream waterfall if's / else's ONLY, not before and / or after these scopes
		if good, error_message := is_even(value); good {
			fmt.Printf("%d is an even number!\n", value)
		} else {
			fmt.Println(error_message)
		}
	}
}
