package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex // Mutex para proteger o acesso ao contador
)

func main() {
	var wg sync.WaitGroup
	const numGoroutines = 10

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(i int) { // Importante passar 'i' como argumento para a função anônima
			defer wg.Done()

			mutex.Lock() // Trava o mutex antes de acessar 'counter'
			x := counter
			x++
			counter = x
			mutex.Unlock() // Destrava o mutex após acessar 'counter'

			fmt.Println("Processo:", i, "| Valor:", x)
		}(i) // Passando 'i' como argumento para garantir o valor correto dentro da goroutine
	}

	wg.Wait()
	fmt.Printf("Valor final do contador: %d\n", counter)
}
