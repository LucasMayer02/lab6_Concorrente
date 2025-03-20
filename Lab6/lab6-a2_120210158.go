package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main()	{
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)

	go func() {
		for i := 0; i < 10000; i++ {
			ch <- rand.Intn(100)
		}
		close(ch)
	}()

	for num := range ch {
		if num%2 == 0 {
			fmt.Println("Par:", num)
		}
	}
}
