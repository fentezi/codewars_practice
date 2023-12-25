package main

import "fmt"

func minOperations(s string) int {
	str := []byte(s)
	index := 0
	last := 0
	count := 0
	for i := 1; i < len(str); i++ {
		if str[index] == str[i] && str[last] == str[i] {
			count++
			str[i] = byte(i)
		}
		index++
		last = index - 1
	}
	return count
}

func main() {
	fmt.Println(minOperations("0100"))
}
