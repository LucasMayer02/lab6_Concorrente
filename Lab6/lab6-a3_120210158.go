package main

import(
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main()	{
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < rand.Intn(5000); i++ {
			ch <- rand.Intn(100)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < rand.Intn(5000); i++ {
			ch <- rand.Intn(100)
		}
	}()

	go func() {
		for num := range ch {
			if num%2 == 0 {
				fmt.Println("Par:", num)
			}
		}
	}()

	wg.Wait()
	close(ch)

}