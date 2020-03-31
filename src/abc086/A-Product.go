package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	if a == 0 {
		fmt.Println("Even")
		return
	} else if b == 0 {
		fmt.Println("Even")
		return
	} else {
		product := a * b

		if product%2 == 0 {
			fmt.Println("Even")
			return
		} else {
			fmt.Println("Odd")
			return
		}
	}
}
