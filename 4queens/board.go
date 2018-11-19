package main

import "math/rand"

const (
	arriba          = 1
	abajo           = 2
	izquierda       = 3
	derecha         = 4
	arribaizquierda = 5
	arribaderecha   = 6
	abajoizquierda  = 7
	abajoderecha    = 8
	dimension       = 4
	operadores      = 8
)

// TEstado estructura para representar un problema
type TEstado struct {
	celdas [dimension][dimension]int
	fila   [dimension * dimension]int
	col    [dimension * dimension]int
}

func shuffle(array []interface{}, source rand.Source) {
	for i := 0; i < len(array)-1; i++ {
		r := rand.Intn(len(array))
		for j := 0; j < len(array)-1; j++ {
			array[i][j], array[j][i] = array[i][j], array[j][i]
		}
	}
}

func estadoInicial() *TEstado {
	return CrearEstado(puzzleInicial)
}

func estadoObjetivo() *TEstado {
	return CrearEstado(puzzleFinal)
}
