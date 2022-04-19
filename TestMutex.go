package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutex   sync.Mutex
	balance int
)

func init() {
	balance = 100
}

func deposit(val int, wg *sync.WaitGroup) {
	mutex.Lock() // lock
	fmt.Println("deposit before", balance)
	balance += val
	fmt.Println("deposit after", balance)
	mutex.Unlock() // unlock
	time.Sleep(time.Millisecond * 10)
	wg.Done()
}

func withdraw(val int, wg *sync.WaitGroup) {
	mutex.Lock() // lock
	fmt.Println("withdraw before", balance)
	balance -= val
	fmt.Println("withdraw after", balance)
	mutex.Unlock() // unlock
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go deposit(20, &wg)
	go withdraw(80, &wg)
	go deposit(40, &wg)
	wg.Wait()
	fmt.Printf("Balance is: %d\n", balance)
}
