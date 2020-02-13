package main

import (
	"fmt"
	"strings"
)

func main() {
	// haiku,atcoder,tasks
	var s string
	fmt.Scan(&s)
	s= strings.Replace(s,",", " ",-1)
	fmt.Println(s)
}
