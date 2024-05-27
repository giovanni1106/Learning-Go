## Resultado

Os testes foram realizados diversas vezes para garantir a integridade do resultado.

    goos: linux
    goarch: amd64
    cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
    BenchmarkAppendPreallocatedSlice
    BenchmarkAppendPreallocatedSlice-8   	1000000000	         0.2479 ns/op
    BenchmarkAppendExistingSlice
    BenchmarkAppendExistingSlice-8       	28570147	        40.52 ns/op
    PASS

## Conclusao

**Prealocação Melhora Performance**: A diferença nos tempos por operação entre os dois benchmarks ilustra claramente o impacto da prealocação. Ao prealocar memória para um slice, você pode evitar realocações custosas, o que resulta em melhorias significativas de performance.