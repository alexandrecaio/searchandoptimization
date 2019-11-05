// Implementação do algoritmo Simm Anneal
package main
 
import (
    "fmt"
    "math"
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
    if (rand.Float64()<0.5){
        return float64(rand.Intn(100))
    }else{
        return float64(-(rand.Intn(100)))
    }
}
 
func printar(solucao Soluction){
    for index := 0; index < 100; index++ {
        fmt.Println("V(",index,")=",solucao.Vetor[index])
    }
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
 
 
func simulatedannealing(solucao Soluction) Soluction{
    t:=1000.0
    var S Soluction
    iteracoes:=1000000
    S = solucao
    var R Soluction
    Best:= S
    for ok := true; ok; ok = iteracoes>0 || t>0{
        R = tweak(S)
        if (qualidade(R)<qualidade(S) || rand.Float64()< math.Exp(( (qualidade(R)-qualidade(S))/t)) ){
            S = R
        }
        t=t-0.1
        if (qualidade(S)<qualidade(Best)){
            Best=S
        }
            
        
        iteracoes-- 
    }
 
    
    
    
    return Best
}
 
 
func qualidade(solucao Soluction) float64{
    var soma float64
    for index := 0; index < 100; index++ {
        soma+= solucao.Vetor[index]*solucao.Vetor[index]
    }
    return soma
}
 
 
func main() {
 
 
    fmt.Println("* Inicio da execução do algoritmo Simulated Annealing.")
    var solucao Soluction
    solucao = gerarSolucaoAleatoria(solucao)
    //printar(solucao)
    
    
 
    fmt.Println("Q(",qualidade(solucao),")")
    solucao = simulatedannealing(solucao)
    fmt.Println("Q(",qualidade(solucao),")")
    
 
 
 
 
}
