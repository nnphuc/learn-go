package main

import (
	"fmt"
)

type Number interface {
	float32 | int
}

func sum[N Number](arr []N) N {
	var result N
	for _, v := range arr {
		result += v
	}
	return result
}

func main() {
	fmt.Println("ok")
	fmt.Println(sum([]int{1, 2, 3}))
}
