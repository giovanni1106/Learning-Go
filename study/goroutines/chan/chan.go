package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Tarefa que os workers irão executar.
// Por simplicidade, nossa "tarefa" é apenas um número que será processado.
type task int

// Resultado de uma tarefa processada.
type result struct {
	task    task
	squared int
}

var wg sync.WaitGroup

// worker é a goroutine que processa tarefas do canal de tarefas e envia os resultados para o canal de resultados.
func worker(id int, tasks <-chan task, results chan<- result) {
	defer wg.Done()
	r := rand.New(rand.NewSource(99))
	for t := range tasks {
		fmt.Printf("Worker %d starting task %d\n", id, t)
		time.Sleep(time.Duration(r.Intn(10)) * time.Second) // Simula um trabalho que leva tempo
		res := result{t, int(t) * int(t)}
		fmt.Printf("Worker %d finished task %d with result %d\n", id, t, res.squared)
		results <- res
	}
}

func main() {
	const numWorkers = 3
	const numTasks = 5

	wg.Add(numWorkers)

	tasks := make(chan task, numTasks)
	results := make(chan result, numTasks)

	// Inicia os workers.
	for i := 1; i <= numWorkers; i++ {
		go worker(i, tasks, results)
	}

	// Envia tarefas para o canal de tarefas.
	for i := 1; i <= numTasks; i++ {
		tasks <- task(i)
	}
	close(tasks) // Importante fechar o canal para indicar que não há mais tarefas a serem enviadas.

	// Coleta os resultados das tarefas.
	for i := 1; i <= numTasks; i++ {
		res := <-results
		fmt.Printf("Task %d result squared is %d\n", res.task, res.squared)
	}

	wg.Wait()

	fmt.Printf("All workers finished\n")
}
