package main

import (
	"fmt"
	"sort"
)

func main() {
	// 7 5 5 OK 17
	// 5 7 5 OK 17
	// 5 5 7 OK 17
	// 5 7 7 NG 19
	// 7 7 3 NG 17

	var intArray = make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&intArray[i])
	}

	sort.Sort(sort.IntSlice(intArray))

	if intArray[0] == 5 && intArray[1] == 5 && intArray[2] == 7 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
