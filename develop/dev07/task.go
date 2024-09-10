package main

import (
	"fmt"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-orChan(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(2*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))

}

func orChan(channels ...<-chan interface{}) <-chan interface{} {
	singleChan := make(chan interface{})
	go func() {
		for {
			for _, ch := range channels {
				select {
				case _, ok := <-ch:
					if !ok {
						close(singleChan)
						return
					}
				default:
					continue
				}
			}
		}
	}()
	return singleChan
}
