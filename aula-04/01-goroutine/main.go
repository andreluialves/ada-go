package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func chamarGoogleSincrono() {
	inicio := time.Now()
	for range 10 {
		resp, err := http.Get("https://www.google.com")
		if err != nil {
			fmt.Println("Erro ao chamar Google:", err)
			return
		}
		defer resp.Body.Close()
	}
	fmt.Println("Tempo total:", time.Since(inicio))
}

func chamarGoogleAssincronoSemWaitGroup() {
	quantidadesDeChamadas := 10
	inicio := time.Now()
	for range quantidadesDeChamadas {
		go func() {
			resp, err := http.Get("https://www.google.com")
			if err != nil {
				fmt.Println("Erro ao chamar Google:", err)
				return
			}
			defer resp.Body.Close()
			fmt.Println("ok")
		}()
	}

	fmt.Println("Tempo total:", time.Since(inicio))
}

func chamarGoogleAssincronoComWaitGroup() {
	var wg sync.WaitGroup
	quantidadesDeChamadas := 10
	wg.Add(quantidadesDeChamadas)
	inicio := time.Now()
	for range quantidadesDeChamadas {
		go func() {
			defer wg.Done()
			resp, err := http.Get("https://www.google.com")
			if err != nil {
				fmt.Println("Erro ao chamar Google:", err)
				return
			}
			defer resp.Body.Close()
			fmt.Println("ok")
		}()
	}
	wg.Wait()
	fmt.Println("Tempo total:", time.Since(inicio))
}

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

type Resultado struct {
	URL    string
	Status string
	Erro   error
}

func chamarGoogleAssincronoComChannel() {
	ch := make(chan Resultado)

	quantidadesDeChamadas := 10
	inicio := time.Now()

	for range quantidadesDeChamadas {
		go func() {
			resp, err := http.Get("https://www.google.com")
			if err != nil {
				ch <- Resultado{
					Status: "erro",
					Erro:   err,
				}
				return
			}
			defer resp.Body.Close()

			ch <- Resultado{
				Status: "ok",
				Erro:   nil,
			}
		}()
	}

	for range quantidadesDeChamadas {
		resultado := <-ch

		if resultado.Erro != nil {
			fmt.Println("Erro ao chamar Google:", resultado.Erro)
			continue
		}

		fmt.Println(resultado.Status)
	}

	fmt.Println("Tempo total:", time.Since(inicio))
}

func chamarGoogleComWaitGroupEChannel() {
	var wg sync.WaitGroup
	ch := make(chan Resultado)

	quantidadesDeChamadas := 5

	for i := 0; i < quantidadesDeChamadas; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			resp, err := http.Get("https://www.google.com")
			if err != nil {
				ch <- Resultado{
					URL:    "https://www.google.com",
					Status: "erro",
					Erro:   err,
				}
				return
			}
			defer resp.Body.Close()

			ch <- Resultado{
				URL:    "https://www.google.com",
				Status: resp.Status,
				Erro:   nil,
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for resultado := range ch {
		if resultado.Erro != nil {
			fmt.Println("Erro ao chamar:", resultado.URL, resultado.Erro)
			continue
		}

		fmt.Println("Resposta:", resultado.URL, resultado.Status)
	}
}

func main() {
	//chamarGoogleSincrono()
	//chamarGoogleAssincronoSemWaitGroup()
	//chamarGoogleAssincronoComWaitGroup()
	//gerarDeadlockComWaitGroup()
	//chamarGoogleAssincronoComChannel()
}
