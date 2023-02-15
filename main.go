package main

import (
	"fmt"
	"sync"
	"time"
)

func factory(start, end int, chan1 chan int, wg *sync.WaitGroup) {
	for i := start; i <= end; i++ {
		time.Sleep(1 * time.Second)
		chan1 <- i
	}
	wg.Done()
}

func consumer(chan1 chan int, chan2 chan bool) {
	for v := range chan1 {
		fmt.Println(v)
		//time.Sleep(1 * time.Second)
	}
	//close(chan1)
	fmt.Println("Hello")
	chan2 <- true
}

func main() {
	factoryNum := 2
	var wg sync.WaitGroup
	start, end := 0, 10
	chan1 := make(chan int)
	for i := 0; i < factoryNum; i++ {
		wg.Add(1)
		go factory(start, end, chan1, &wg)
		start += 10
		end += 10
	}
	consumerDone := make(chan bool)
	go consumer(chan1, consumerDone)
	wg.Wait()
	close(chan1)
	<-consumerDone
	fmt.Println("Finished all goroutines")
}
