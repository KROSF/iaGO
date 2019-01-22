package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	busquedaInformada(aestrella)
	fmt.Println("Busqueda en A*: ", time.Since(start))
	// inicial := estadoInicial()
	// operador := 0
	// for !inicial.testObjetivo() {
	// 	if inicial.esValido(operador) {
	// 		inicial = inicial.aplicaOperador(operador)
	// 	}
	// 	inicial.dispEstado()
	// 	fmt.Scanf("%d", &operador)
	// }
}
