package main

import (
	"fmt"
	"sync"
	"time"
)

var balance1 int

func init() {
	balance1 = 100
}

func deposit1(val int, wg sync.WaitGroup, ch chan bool) {
	ch <- true
	fmt.Println("deposit before", balance1)
	balance1 += val
	fmt.Println("deposit after", balance1)
	<-ch
	time.Sleep(time.Millisecond * 10)
	wg.Done()
}

func withdraw1(val int, wg sync.WaitGroup, ch chan bool) {
	ch <- true
	fmt.Println("withdraw before", balance1)
	balance1 -= val
	fmt.Println("withdraw after", balance1)
	<-ch
	wg.Done()
}

func main() {
	ch := make(chan bool, 1)
	var wg sync.WaitGroup
	go deposit1(20, wg, ch)
	go withdraw1(80, wg, ch)
	go deposit1(40, wg, ch)
	<-ch
	fmt.Printf("Balance is: %d\n", balance)
}
