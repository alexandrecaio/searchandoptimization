// Implementação do algoritmo Hill CLimbing
package main
 
import (
    "fmt"
    "math/rand"
    
)
 
type Soluction struct {
    Vetor [100]float64
}
 
func gerarSolucaoAleatoria(solucao Soluction) Soluction{
    for index := 0; index < 100; index++ {
        solucao.Vetor[index]=float64(rand.Intn(100))
    }
    return solucao
}
 
func gerarNumeroAleatorio() float64{
    return float64(rand.Intn(100))
}
 
 
 
func tweak(solucao Soluction) Soluction{
    p:=1.0
    //r:=100.0 range
    min:=0.0
    max:=100.0
    n:=0.0
    for index := 0; index < 100; index++ {
        if(p>rand.Float64()){       
            for ok := true; ok; ok = (solucao.Vetor[index]+n )>min && (solucao.Vetor[index]+n )<max {
                n=gerarNumeroAleatorio()
            }
            solucao.Vetor[index]=n
            
        }
    }
 
    return solucao
}
 
 
func hillclimbing(solucao Soluction) Soluction{
    var S Soluction
    iteracoes:=100000
    S = solucao
    var R Soluction
    for ok := true; ok; ok = iteracoes>0 || qualidade(S)==0 {
        R = tweak(S)
        if (qualidade(R)<qualidade(S)){
            S = R
        }
        iteracoes-- 
    }
    
    return S
}
 
 
func qualidade(solucao Soluction) float64{
    var soma float64
    for index := 0; index < 100; index++ {
        soma+= solucao.Vetor[index]*solucao.Vetor[index]
    }
    return soma
}
 
 
func main() {
 
    fmt.Println("* Inicio da execução do algoritmo Hill Climbing.")
    var solucao Soluction
    solucao = gerarSolucaoAleatoria(solucao)
 
    fmt.Println("Q(",qualidade(solucao),")")
    solucao = hillclimbing(solucao)
    fmt.Println("Q(",qualidade(solucao),")")
    
 
 
 
 
}
