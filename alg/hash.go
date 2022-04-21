package main

import "fmt"

func main() {
	fmt.Println(Try(10, 3))
}

var m map[[2]int]int64

func Try(n int, k int) int64 {
	if m == nil {
		m = make(map[[2]int]int64)
	}
	if k < 0 {
		return 0
	}
	pair := [2]int{n, k}
	if v, ok := m[pair]; ok {
		return v
	}
	if n == 1 {
		if k >= 0 && k <= 9 {
			return 1
		} else {
			return 0
		}
	} else {
		var res int64
		for i := 0; i < 10; i++ {
			res += Try(n-1, k-i)
		}
		//fmt.Println(pair)
		m[pair] = res
		return res
	}
}
