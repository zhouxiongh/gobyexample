package main

import (
	"fmt"
	"sync"
	"time"
)

type Raft struct {
	mu sync.Mutex
	cond *sync.Cond
	b bool
}

func (rf *Raft) applier() {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	for i := 0; i < 5; i++{
		if rf.b {
			fmt.Println("cond changed")
		} else {
			fmt.Println("wait")
			rf.cond.Wait()
			fmt.Println("wake")
		}
	}
}

func main() {
	rf := &Raft{}
	rf.cond = sync.NewCond(&rf.mu)
	go rf.applier()
	time.Sleep(time.Second)
	rf.b = true
	rf.cond.Signal()
	time.Sleep(time.Second)
}
