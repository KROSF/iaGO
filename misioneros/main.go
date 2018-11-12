package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	busquedaInformada(aestrella)
	fmt.Println("Tiempo hasta la solucion: ", time.Since(start))
	// e := estadoInicial()
	// for e.esObjetivo() != true {
	// 	e.dispEstado()
	// 	op := 0
	// 	fmt.Scanf("%d", &op)
	// 	if e.esValido(op) {
	// 		e = e.aplicaOperador(op)
	// 	}
	// }
}
