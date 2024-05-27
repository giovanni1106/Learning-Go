package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex1, mutex2 sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1
	go func() {
		mutex1.Lock()
		fmt.Println("Goroutine 1 pegou o Mutex 1")
		time.Sleep(time.Second)
		mutex2.Lock()
		fmt.Println("Goroutine 1 pegou o Mutex 2")
		time.Sleep(time.Second)

		// Liberando os mutexes
		mutex2.Unlock()
		mutex1.Unlock()

		wg.Done()
	}()

	// Goroutine 2
	go func() {
		mutex1.Lock()
		fmt.Println("Goroutine 2 pegou o Mutex 1")
		time.Sleep(time.Second)
		mutex2.Lock()
		fmt.Println("Goroutine 2 pegou o Mutex 2")
		time.Sleep(time.Second)

		// Liberando os mutexes
		mutex2.Unlock()
		mutex1.Unlock()

		wg.Done()
	}()

	wg.Wait()
}
