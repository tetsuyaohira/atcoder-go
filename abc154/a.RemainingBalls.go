package main

import "fmt"

/**
Problem Statement
文字列Sの書かれたボールがA個、文字列Tの書かれたボールがB個あります。
高橋君は、文字列Uの書かれたボールを1個選んで捨てました。
文字列S,Tの書かれたボールがそれぞれ何個残っているか求めてください

Sample input1
red blue
3 4
red
*/
func main() {

	var s, t, u string
	var sCnt, tCnt int

	fmt.Scan(&s, &t)
	fmt.Scan(&sCnt, &tCnt)
	fmt.Scan(&u)

	if s == u {
		sCnt -= 1
	}
	if t == u {
		tCnt -= 1
	}

	fmt.Println(sCnt, tCnt)
}
