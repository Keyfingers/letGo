package main

import (
	"fmt"
	"sync"
)

type MyMutex struct {
	count int
	sync.Mutex
}

func main() {
	var mu MyMutex
	mu.Lock()
	var mu2 = mu //复制锁，虽然会复制出两个独立的锁，但是锁的状态也会被复制，所以会mu已经lock，mu2初始也会是lock，再执mu2.lock会造成死锁
	mu.count++
	mu.Unlock()
	mu2.Lock()
	mu2.count++
	mu2.Unlock()
	fmt.Println(mu.count, mu2.count)
}
