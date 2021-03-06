package main

import(
    "fmt"
    "time"
)

func generarNumeros(out chan<- int) {
    for x := 0; x < 5; x++ {
        out <- x
    }
    close(out)
}

func elevarAlCuadrado(in <-chan int, out chan int) {
    for x := range in {
        out <- x * x
    }
    close(out)
}

func imprimir(in <-chan int) {
    for x := range in {
        fmt.Println(x)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    numero := make(chan int)
    cuadrado := make(chan int)

    go generarNumeros(numero)
    go elevarAlCuadrado(numero, cuadrado)

    imprimir(cuadrado)
}
