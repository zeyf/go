package main

import (
	"fmt"
)

/*
An even-ended number is a number with the same first and last digits
Count all the even-ended numbers that result from multiplying two four-digit numbers
*/

func is_even_ended_number(number_as_string string) bool {
	return number_as_string[0] == number_as_string[len(number_as_string)-1]
}

func main() {
	var even_ended_numbers_count int = 0

	for x := 1000; x < 10000; x++ {
		for y := x; y < 10000; y++ {
			product := x * y
			product_string := fmt.Sprintf("%d", product)

			if is_even_ended_number(product_string) {
				even_ended_numbers_count += 1
			}
		}
	}

	fmt.Printf("Found %d even-ended numbers!\n", even_ended_numbers_count)
}
