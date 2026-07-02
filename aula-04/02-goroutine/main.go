package main

import (
	"fmt"
	"sync"
)

func gerarDeadlockComWaitGroup() {
	var wg sync.WaitGroup

	wg.Add(10) // estou dizendo: vou esperar 10 Done()

	for i := 0; i < 5; i++ { // mas crio só 5 goroutines
		go func() {
			defer wg.Done()
			fmt.Println("ok")
		}()
	}

	wg.Wait() // vai esperar mais 5 Done() que nunca virão

	fmt.Println("Isso nunca será executado")
}

func main() {
	gerarDeadlockComWaitGroup()
}
