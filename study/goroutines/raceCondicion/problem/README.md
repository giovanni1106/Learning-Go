## Resultado

O resultado da execução do programa demonstra claramente uma condição de corrida, onde o acesso concorrente ao recurso compartilhado (counter) resulta em uma atualização inconsistente do valor da variável. Notavelmente, o valor final de counter não corresponde à quantidade de incrementos realizados pelas goroutines, ilustrando a perda de atualizações devido à sobreposição de leituras e escritas.

    Processo: 9 | Valor: 1
    Processo: 3 | Valor: 2
    Processo: 4 | Valor: 3
    Processo: 5 | Valor: 4
    Processo: 6 | Valor: 5
    Processo: 7 | Valor: 6
    Processo: 8 | Valor: 7
    Processo: 0 | Valor: 1
    Processo: 2 | Valor: 1
    Processo: 1 | Valor: 2
    Valor final do contador: 2

## Conclusão

O comportamento observado destaca a importância de mecanismos de sincronização ao lidar com acessos concorrentes a recursos compartilhados. Sem esses mecanismos, os programas estão sujeitos a condições de corrida que podem levar a dados corrompidos e comportamentos indeterminados.