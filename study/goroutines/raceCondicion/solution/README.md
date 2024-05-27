# Solução para Condição de Corrida em Goroutines

Este diretório contém a solução para o problema de condição de corrida observado no acesso concorrente a uma variável compartilhada entre várias goroutines. A condição de corrida foi resolvida usando `sync.Mutex` para sincronizar o acesso à variável `counter`.

## Problema Original

O problema inicial demonstrou um cenário onde múltiplas goroutines incrementavam uma variável compartilhada (`counter`) sem sincronização apropriada, resultando em uma atualização inconsistente de seu valor.

## Solução Implementada

Para resolver a condição de corrida, utilizamos um `sync.Mutex` para garantir que apenas uma goroutine por vez pudesse modificar o valor de `counter`. Isso é feito envolvendo as operações de leitura, incremento e escrita de `counter` com chamadas para `mutex.Lock()` e `mutex.Unlock()`.

```go
mutex.Lock() // Trava o mutex antes de acessar 'counter'
x := counter
x++
counter = x
mutex.Unlock() // Destrava o mutex após acessar 'counter'
```

## Resultado

A aplicação da solução resultou em um comportamento consistente e esperado, com cada goroutine incrementando a variável compartilhada de forma segura, e o programa alcançando o valor final correto para counter.


    Processo: 2 | Valor: 1
    Processo: 3 | Valor: 4
    Processo: 5 | Valor: 6
    Processo: 1 | Valor: 10
    Processo: 4 | Valor: 5
    Processo: 9 | Valor: 2
    Processo: 8 | Valor: 9
    Processo: 6 | Valor: 7
    Processo: 7 | Valor: 8
    Processo: 0 | Valor: 3
    Valor final do contador: 10

## Conslusão

A utilização de sync.Mutex provou ser uma estratégia eficaz para resolver condições de corrida em programas Go que utilizam concorrência. Este exemplo sublinha a importância da sincronização adequada ao acessar recursos compartilhados em ambientes concorrentes para prevenir comportamentos indesejados e garantir a integridade dos dados.