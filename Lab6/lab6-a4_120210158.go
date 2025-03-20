package main

import(
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producer(ch chan <- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5000; i++ {
		ch <- rand.Intn(100)
	}
}

func consumer(ch <-chan int) {
	for num := range ch {
		if num%2 == 0 {
			fmt.Println("Par", num)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int, 100)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(ch, &wg)
	go producer(ch, &wg)

	go consumer(ch)

	wg.Wait()
	close(ch)
}