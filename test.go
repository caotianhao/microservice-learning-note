package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var i int64 = 10
	k := &i
	var mu sync.Mutex

	ticker := time.NewTicker(time.Second / 10)
	defer ticker.Stop()

	done := make(chan struct{})

	go func() {
		time.Sleep(time.Second)
		done <- struct{}{}
	}()

	time.Sleep(5 * time.Second)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Done!", i)
				return
			case <-ticker.C:
				mu.Lock()
				atomic.AddInt64(k, int64(1))
				mu.Unlock()
			}
		}
	}()
}
