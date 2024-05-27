## Resultados

A saída do programa mostra claramente a intercalação na execução das goroutines. Mesmo que uma goroutine comece a executar seu loop, ela pode ser pausada para que outra goroutine inicie ou continue sua execução, demonstrando o trabalho do escalonador de goroutines em ação.


    Esperando as goroutines terminarem
    0: 0
    4: 0
    0: 1
    4: 1
    .
    .
    .
    Goroutine 4 completou
    1: 49
    Goroutine 1 completou
    Todas as goroutines terminaram


## Conclusão

A execução intercalada das goroutines evidencia o design concorrente do Go, permitindo que várias tarefas progridam de forma eficiente sem uma espera sequencial rígida para que cada uma complete. Este comportamento maximiza o uso do processador, especialmente em sistemas com múltiplos núcleos, e destaca a eficácia do Go na gestão de concorrência leve com goroutines.