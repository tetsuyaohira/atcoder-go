package main

import "fmt"

func main() {
	// in  9 45000
	// out 4 0 5
	// 10000 5000 1000

	var n, yen int
	fmt.Scan(&n)
	fmt.Scan(&yen)

	fmt.Println(calcYen(n, yen))
	calcYen(n, yen)

}

func calcYen(n int, yen int) (int, int, int) {
	for a := 0; a <= n; a++ {
		for b := 0; b <= n; b++ {

			if n-a-b < 0 {
				continue
			}

			if a*10000+b*5000+(n-a-b)*1000 == yen {
				return a, b, n - a - b
			}
		}
	}
	return -1, -1, -1
}
