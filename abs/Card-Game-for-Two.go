package main

import (
  "fmt"
  "sort"
)

//3
//2 7 4
func main() {

  var n int
  fmt.Scan(&n)

  var cards = make([]int, n)
  for i:=0 ; i<n ; i++ {
    fmt.Scan(&cards[i])
  }
  sort.Sort(sort.Reverse(sort.IntSlice(cards)))

  var result int
  for i:=0 ;i <n ; i++ {
    if (i+1)%2 == 1 {
      result += cards[i]
    } else {
      result -= cards[i]
    }
  }

  fmt.Println(result)

}

