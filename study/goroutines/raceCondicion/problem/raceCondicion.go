package main

import (
	"fmt"
	"sync"
)

// Contador compartilhado por várias goroutines
var counter int

func main() {
	var wg sync.WaitGroup
	const numGoroutines = 10

	wg.Add(numGoroutines)

	// Incrementa o contador em 1, 100 vezes, em goroutines separadas
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()

			x := counter
			x++

			fmt.Println("Processo:", i, "| Valor:", x)

			counter = x
		}(i)
	}

	wg.Wait()
	fmt.Printf("Valor final do contador: %d\n", counter)
}
