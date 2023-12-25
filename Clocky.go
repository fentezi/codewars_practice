package main

import "fmt"

func WhatTimeIsIt(angle int) string {
	var time int
	switch angle {
	case 0:
		return "12:00"
	case 90:
		return "03:00"
	case 180:
		return "06:00"
	case 270:
		return "09:00"
	case 360:
		return "12:00"
	default:
		for angle > 0 {
			time += 2
			angle -= 1
		}
	}
	hours := time / 60
	if hours == 0 {
		return fmt.Sprintf("12:%02d", time%60)
	}
	return fmt.Sprintf("%02d:%02d", time/60, time%60)
}

func main() {
	fmt.Println(WhatTimeIsIt(10))
}
