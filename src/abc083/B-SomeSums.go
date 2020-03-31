package main

import "fmt"

func main() {
	var n, min, max int
	fmt.Scanf("%d %d %d", &n, &min, &max)

	var sum int
	for i := 1; i <= n; i++ {
		tmp := i
		var remainder int
		for tmp > 0 {
			remainder += tmp % 10
			tmp /= 10
		}

		if remainder >= min && remainder <= max {
			sum += i
		}
	}

	fmt.Println(sum)
}
