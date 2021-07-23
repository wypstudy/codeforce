package main

import "fmt"

func main() {
	var (
		t int
		s string
	)
	fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		fmt.Scanln(&s)
		red, green := make(map[rune]bool), make(map[rune]bool)
		for _, ch := range s {
			if !red[ch] {
				red[ch] = true
			} else if red[ch] {
				green[ch] = true
			}
		}
		redSize, greenSize := len(red), len(green)
		fmt.Println(greenSize + (redSize-greenSize)/2)
	}
}
