package main

import (
	"fmt"
	"sync"
	"time"
)

func print(number int) {
	for i := 0; i < 10; i++ {
		fmt.Println(number)
	}

}

func main() {
	var wg sync.WaitGroup
	n := time.Now()
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			print(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("processed in", time.Since(n))
}
