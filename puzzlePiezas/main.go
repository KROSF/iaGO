package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	busquedaInformada(vorax)
	fmt.Println("Busqueda en A*: ", time.Since(start))
	//start = time.Now()
	//busquedaAnchura()
	//fmt.Println("Busqueda Anchura: ", time.Since(start))
}
