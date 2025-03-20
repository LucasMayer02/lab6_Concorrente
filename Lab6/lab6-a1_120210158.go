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
		for {
			ch <- rand.Intn(100)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	for num := range ch {
		if num%2 == 0 {
			fmt.Println("Par:", num)
		}
	}
}