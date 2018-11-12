package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	busquedaInformada(vorax)
	fmt.Println("Tiempo hasta la solucion: ", time.Since(start))
}
