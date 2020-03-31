package main

import (
  "fmt"
)

/**
in
2
2
2
100
 */
func main() {
  var a, b, c, target int
  fmt.Scan(&a, &b, &c, &target)
  cnt := 0
  for i:=0 ; i <= a ; i++ {
    for j:=0 ; j<=b ; j++ {
      for k:=0 ; k<=c ; k++ {
        if i*500 + j*100 + k*50 == target {
          cnt++
        }
      }
    }
  }
  fmt.Println(cnt)
}

