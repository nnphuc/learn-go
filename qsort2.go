package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Data interface {
	Less(i, j int) bool
	Len() int
	Swap(i, j int)
}

func pivot(data Data, left, right int) int {
	if left >= right {
		return left
	}

	pivot := left
	//t := data.(Array)
	//pivotValue := t.A[right-1]

	for i := left; i < right-1; i++ {
		if data.Less(i, right-1) {
			data.Swap(i, pivot)
			pivot++
		}
	}
	data.Swap(pivot, right-1)
	//fmt.Println("pivot", left, right, pivot, t, pivotValue)

	return pivot
}

func qsortUtil(data Data, left, right int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	//fmt.Println("qSortUtil", left, right, data)
	if left >= right {
		return
	}
	p := pivot(data, left, right)

	//wg2 := sync.WaitGroup{}
	//wg2.Add(2)
	//go qsortUtil(data, left, p-1, &wg2)
	//go qsortUtil(data, p+1, right, &wg2)
	//wg2.Wait()

	go qsortUtil(data, left, p-1, wg)
	go qsortUtil(data, p+1, right, wg)
}

func qsort(data Data) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	qsortUtil(data, 0, data.Len(), &wg)
	wg.Wait()
}

type Array struct {
	A []int
}

func (a Array) Len() int {
	return len(a.A)
}

func (a Array) Less(i, j int) bool {
	return a.A[i] < a.A[j]
}

func (a Array) Swap(i, j int) {
	t := a.A[i]
	a.A[i] = a.A[j]
	a.A[j] = t
}

func main() {
	var arr []int
	for i := 0; i < 10_000_000; i++ {
		arr = append(arr, i)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	a := Array{
		A: arr,
	}
	start := time.Now()
	qsort(a)
	end := time.Now()

	fmt.Println(end.Sub(start).Microseconds())
	//fmt.Println(a.A)
}
