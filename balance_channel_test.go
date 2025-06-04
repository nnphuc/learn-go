package main

import (
	"sync"
	"testing"
)

func TestBalanceWithChannel(t *testing.T) {
	balance1 = 100
	ch := make(chan bool, 1)
	var wg sync.WaitGroup
	wg.Add(3)
	go deposit1(20, &wg, ch)
	go withdraw1(80, &wg, ch)
	go deposit1(40, &wg, ch)
	wg.Wait()
	if balance1 != 80 {
		t.Fatalf("expected balance 80, got %d", balance1)
	}
}
