package main

import (
	"context"
	"fmt"
	"time"
)

func pegarValorDoChannel(ch chan<- string, duration time.Duration) {
	time.Sleep(duration)
	ch <- "Acabou"
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	ch := make(chan string)
	go pegarValorDoChannel(ch, 1*time.Second)

	ch2 := make(chan string)
	go pegarValorDoChannel(ch2, 3*time.Second)

	ch3 := make(chan string)
	go pegarValorDoChannel(ch3, 4*time.Second)

	for {
		select {
		case returnValue := <-ch:
			fmt.Println(returnValue)
		case returnValue := <-ch2:
			fmt.Println(returnValue)
		case returnValue := <-ch3:
			fmt.Println(returnValue)
		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		}
	}
}
