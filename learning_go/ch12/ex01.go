package main

import (
	"fmt"
	"sync"
)

func ch12ex01() {
	ch := make(chan int)
	var wg0 sync.WaitGroup
	wg0.Add(2)
	go func() {
		defer wg0.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		defer wg0.Done()
		for i := 11; i < 20; i++ {
			ch <- i
		}
	}()
	go func() {
		wg0.Wait()
		close(ch)
	}()
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		for i := range ch {
			fmt.Println(i)
		}
	}()
	wg1.Wait()
}
