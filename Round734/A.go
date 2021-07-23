package main

import "fmt"

func main() {
	var (
		t int
		n int
	)
	fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		fmt.Scanln(&n)
		c := n / 3
		c1, c2 := c, c
		if n%3 == 1 {
			c1 += 1
		}
		if n%3 == 2 {
			c2 += 1
		}
		fmt.Println(c1, c2)
	}
}
