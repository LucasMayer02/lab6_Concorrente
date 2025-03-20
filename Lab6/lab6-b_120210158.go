package main

import(
	"fmt"
	"math/rand"
	"sync"
	"time"
)


func request() int {
	n := rand.Intn(10)
	fmt.Printf("Goroutine sorteou: %d\n", n)
	time.Sleep(time.Duration(n) * time.Second)
	return n
}

func gateway(ngo, wait_n int) int {
	rand.Seed(time.Now().UnixNano())

	results := make(chan int, ngo)
	var wg sync.WaitGroup
	var sum int

	for i := 0; i < ngo; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n := request()
			results <- n
		}()
	}

	go func() {
		for i := 0; i < wait_n; i++ {
			sum += <- results
		}
		close(results)
	}()

	wg.Wait()

	return sum
}

func main() {
	ngo := 5
	wait_n := 3

	fmt.Printf("\nSoma dos valores: %d\n", gateway(ngo, wait_n))
}