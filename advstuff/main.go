package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	dataChan := make(chan int)

	go func() {
		Wg := sync.WaitGroup{}
		for i := 0; i < 10000; i++ {
			Wg.Add(1)
			go func() {
				defer Wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}
		Wg.Wait()
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Printf("n = %d\n", n)
	}
}

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}
