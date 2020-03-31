package main

import (
	"fmt"
	"sort"
)

func main() {

	cnt := 0
	fmt.Scan(&cnt)

	var kagamis = make([]int, cnt)
	for i := 0; i < cnt; i++ {
		fmt.Scan(&kagamis[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(kagamis)))

	result := 0
	var prevKagami = 101
	for i := 0; i < cnt; i++ {
		if kagamis[i] < prevKagami {
			result++
		}
		prevKagami = kagamis[i]
	}
	fmt.Println(result)
}
