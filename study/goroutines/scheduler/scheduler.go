package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup // Inicializa o WaitGroup

	// Define o número de goroutines
	numGoroutines := 5
	wg.Add(numGoroutines) // Incrementa o contador do WaitGroup

	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()

			for j := 0; j < 50; j++ {
				fmt.Printf("%v: %v\n", i, j)
			}

			fmt.Println("Goroutine", i, "completou")
		}(i)
	}

	fmt.Println("Esperando as goroutines terminarem")
	wg.Wait()
	fmt.Println("Todas as goroutines terminaram")
}
