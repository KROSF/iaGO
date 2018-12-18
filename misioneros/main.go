package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	busquedaInformada(aestrella)
	fmt.Println("Tiempo hasta la solucion: ", time.Since(start))
}
