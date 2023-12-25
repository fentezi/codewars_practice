package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EncryptThis(text string) string {
	if text == "" {
		return ""
	}
	var result []string
	for _, t := range strings.Split(text, " ") {
		ans := ""
		length := len(t)
		ans += strconv.Itoa(int(t[0]))
		if length > 2 {
			scd, lst := t[length-1], t[1]
			ans += string(scd)
			ans += t[2 : length-1]
			ans += string(lst)
		} else {
			ans += t[1:]
		}
		result = append(result, ans)
	}
	return strings.Join(result, " ")
}

func main() {
	fmt.Println(EncryptThis(""))
}
