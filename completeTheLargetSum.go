package main

import "fmt"

func LargestSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	result := 0
	currentSum := 0

	for _, value := range arr {
		currentSum = maxx(value, currentSum+value)
		result = maxx(result, currentSum)
	}

	if result < 0 {
		return 0
	}

	return result
}

func maxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{31, -41, 59, 26, -53, 58, 97, -93, -23, 84}
	fmt.Println(LargestSum(nums))
}
