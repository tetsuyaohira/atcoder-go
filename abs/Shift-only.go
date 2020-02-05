package main

import "fmt"

func main() {

  var cnt int
  fmt.Scan(&cnt);
  
  var nums = make([]int,cnt)
  for i := 0 ; i < cnt ; i++ {
     fmt.Scan(&nums[i])
  }
  
  var flag = true
  var result = 0
  for flag {
    for i :=0 ; i< cnt ; i++ {
      if nums[i]%2 == 0 {
        nums[i] = nums[i] / 2
      } else {
      	flag = false
      }
    }
    if flag == true {
      result = result +1 
    }
  }
  
  fmt.Println(result)
}

