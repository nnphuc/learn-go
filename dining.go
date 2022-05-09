package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Fork struct {
	id   int
	lock sync.Mutex
}

type Philosopher struct {
	name  string
	left  *Fork
	right *Fork
}

var count int32

func (p *Philosopher) eat() {
	if atomic.LoadInt32(&count) == 0 {
		return
	}
	atomic.AddInt32(&count, -1)
	defer atomic.AddInt32(&count, 1)
	canUseLeft := p.left.lock.TryLock()
	if !canUseLeft {
		return
	}
	defer p.left.lock.Unlock()

	canUseRight := p.right.lock.TryLock()
	if !canUseRight {
		return
	}
	defer p.right.lock.Unlock()
	fmt.Printf("%s is start eating\n", p.name)
	//time.Sleep(time.Millisecond * 100)
	fmt.Printf("%s is done eating\n", p.name)
	return
}

func main() {
	num := 5
	forks := []Fork{}
	for i := 0; i < num; i++ {
		f := Fork{
			id:   i,
			lock: sync.Mutex{},
		}
		forks = append(forks, f)
	}

	ps := []Philosopher{}

	for i := 0; i < num; i++ {
		p := Philosopher{
			name:  fmt.Sprintf("Philosopher %d", i),
			left:  &forks[i],
			right: &forks[(i+1)%len(forks)],
		}
		ps = append(ps, p)
	}

	count = 2
	for i := 0; i < num; i++ {
		go func(i int) {
			for {
				time.Sleep(time.Millisecond * 30)
				ps[i].eat()
			}
		}(i)
	}
	time.Sleep(time.Minute * 10)
}
