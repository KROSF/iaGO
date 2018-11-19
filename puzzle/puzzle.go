package main

import (
	"fmt"
	"math"
)

const (
	arriba     = 1
	abajo      = 2
	izquierda  = 3
	derecha    = 4
	dimension  = 3
	operadores = 4
)

// TEstado estructura para representar un problema
type TEstado struct {
	celdas [dimension][dimension]int
	fila   [dimension * dimension]int
	col    [dimension * dimension]int
}

var puzzleInicial = [][]int{{0, 2, 3}, {1, 8, 4}, {7, 6, 5}}
var puzzleFinal = [][]int{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}

//CrearEstado devulve un TEstado creado desde un slice
func CrearEstado(puzzle [][]int) *TEstado {
	e := TEstado{}
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[0]); j++ {
			ficha := puzzle[i][j]
			e.celdas[i][j] = ficha
			e.fila[ficha] = i
			e.col[ficha] = i
		}
	}
	return &e
}

func estadoInicial() *TEstado {
	return CrearEstado(puzzleInicial)
}

func estadoObjetivo() *TEstado {
	return CrearEstado(puzzleFinal)
}

func (e *TEstado) iguales(f *TEstado) bool {
	for i := 0; i < len(e.celdas); i++ {
		for j := 0; j < len(e.celdas[0]); j++ {
			if e.celdas[i][j] != f.celdas[i][j] {
				return false
			}
		}
	}
	return true
}

func (e *TEstado) testObjetivo() bool {
	return e.iguales(estadoObjetivo())
}

func (e *TEstado) esValido(op int) bool {
	valido := false
	switch op {
	case arriba:
		valido = (e.fila[0] > 0)
		break
	case abajo:
		valido = (e.fila[0] < len(e.celdas)-1)
		break
	case izquierda:
		valido = (e.col[0] > 0)
		break
	case derecha:
		valido = (e.col[0] < len(e.celdas)-1)
		break
	}
	return valido
}

func (e *TEstado) intercambio(fnew, cnew, fold, cold int) {
	ficha := e.celdas[fnew][cnew]
	e.celdas[fold][cold] = ficha
	e.celdas[fnew][cnew] = 0
	e.col[ficha] = cold
	e.fila[ficha] = fold
}

func (e *TEstado) aplicaOperador(op int) *TEstado {
	nuevo := *e
	switch op {
	case arriba:
		nuevo.fila[0]--
		break
	case abajo:
		nuevo.fila[0]++
		break
	case izquierda:
		nuevo.col[0]--
		break
	case derecha:
		nuevo.col[0]++
		break
	}
	nuevo.intercambio(nuevo.fila[0], nuevo.col[0], e.fila[0], e.col[0])
	return &nuevo
}

func (e *TEstado) coste(op int) int {
	return 1
}

func (e *TEstado) dispEstado() {
	for i := 0; i < len(e.celdas); i++ {
		fmt.Println(e.celdas[i])
	}
	fmt.Println()
}

func dispOperador(op int) {
	switch op {
	case arriba:
		fmt.Printf("ARRIBA\n\n")
		break
	case abajo:
		fmt.Printf("ABAJO\n\n")
		break
	case izquierda:
		fmt.Printf("IZQUIERDA\n\n")
		break
	case derecha:
		fmt.Printf("DERECHA\n\n")
		break
	}
}

func (e *TEstado) heuristica() int {
	return e.manhattan()
}

func (e *TEstado) manhattan() int {
	obj := estadoObjetivo()
	sum := 0.0
	for i := 0; i < len(e.col); i++ {
		sum += math.Abs(float64(e.fila[i]-obj.fila[i])) + math.Abs(float64(e.col[i]-obj.col[i]))
	}
	return int(sum)
}

func (e *TEstado) piezasMalColodas() int {
	sum := 0
	obj := estadoObjetivo()
	for i := 0; i < len(e.celdas); i++ {
		for j := 0; j < len(e.celdas[0]); j++ {
			if e.celdas[i][j] != obj.celdas[i][j] {
				sum++
			}
		}
	}
	return sum
}
