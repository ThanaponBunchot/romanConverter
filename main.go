package main

import (
	"fmt"
)

func main() {
	// 	"MM"      -> 2000
	// "MDCLXVI" -> 1666
	// "M"       -> 1000
	// "CD"      ->  400
	// "XC"      ->   90
	// "XL"      ->   40
	// "I"       ->    1

	// Symbol    Value
	// I          1
	// V          5
	// X          10
	// L          50
	// C          100
	// D          500
	// M          1,000

	res := Decode("MCXXIX")
	fmt.Printf("res : %d \n", res)

	// 2919 from roman MMCMXIX
}
func Decode(roman string) int {

	number := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	arr := []int{}
	for _, v := range roman {
		arr = append(arr, number[string(v)])
	}
	result := 0
	prev := 0

	for i := len(arr) - 1; i >= 0; i-- {
		value := arr[i]
		if prev > value {
			result -= value
		} else {
			result += value
		}
		prev = value
	}
	return result
}
