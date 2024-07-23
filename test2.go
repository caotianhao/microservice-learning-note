package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		i  int
		mu sync.Mutex
		wg sync.WaitGroup
	)

	done := make(chan struct{})

	for j := 0; j < 100; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			i++
			mu.Unlock()
		}()
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	<-done

	fmt.Println(i)
}
