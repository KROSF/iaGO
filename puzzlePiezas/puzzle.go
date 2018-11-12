package main

import (
	"fmt"
	"math"
)

const (
	arribaA    = 1
	abajoA     = 2
	izquierdaA = 3
	derechaA   = 4
	arribaB    = 5
	abajoB     = 6
	izquierdaB = 7
	derechaB   = 8
	arribaC    = 9
	abajoC     = 10
	izquierdaC = 11
	derechaC   = 12
	dimension  = 6
	operadores = 12
	vacia      = 0
	obstaculo  = -1
	a          = 1
	b          = 2
	c          = 3
)

type tEstado struct {
	celdas [dimension][dimension]int
	pos    [2][dimension + 1]int
}

var puzzleInicial = [][]int{
	{-1, 0, 0, 3, 0, 0},
	{-1, 0, 0, 3, 0, 0},
	{0, 1, 0, 3, 0, 0},
	{1, 1, 1, -1, 2, 0},
	{0, 1, 0, 2, 2, 2},
	{0, 0, 0, 0, 0, 0}}
var puzzleFinal = [][]int{
	{-1, 0, 0, 0, 0, 0},
	{-1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0},
	{0, 1, 0, -1, 0, 3},
	{1, 1, 1, 2, 0, 3},
	{0, 1, 2, 2, 2, 3}}

func crearEstado(puzzle [][]int) *tEstado {
	e := tEstado{}
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[0]); j++ {
			ficha := puzzle[i][j]
			e.celdas[i][j] = ficha
			f, c := esCentro(ficha, i, j, puzzle)
			if f != -1 && c != -1 {
				e.pos[0][ficha] = f
				e.pos[1][ficha] = c
			}
		}
	}
	return &e
}

func esCentro(ficha, i, j int, p [][]int) (int, int) {
	switch ficha {
	case a:
		if i > 0 && i < dimension-1 && j > 0 && j < dimension-1 {
			if p[i][j-1] == a && p[i][j+1] == a && p[i-1][j] == a && p[i+1][j] == a {
				return i, j
			}
		}
		break
	case b:
		if i > 0 && i < dimension && j > 0 && j < dimension-1 {
			if p[i][j-1] == b && p[i][j+1] == b && p[i-1][j] == b {
				return i, j
			}
		}
		break
	case c:
		if i > 0 && i < dimension-1 && j >= 0 && j < dimension {
			if p[i-1][j] == c && p[i+1][j] == c {
				return i, j
			}
		}
		break
	}
	return -1, -1
}

func estadoInicial() *tEstado {
	return crearEstado(puzzleInicial)
}

func estadoObjetivo() *tEstado {
	return crearEstado(puzzleFinal)
}

func (e *tEstado) iguales(f *tEstado) bool {
	for i := 0; i < len(e.pos); i++ {
		for j := 0; j < len(e.pos[0]); j++ {
			if e.pos[i][j] != f.pos[i][j] {
				return false
			}
		}
	}
	return true
}

func (e *tEstado) testObjetivo() bool {
	return e.iguales(estadoObjetivo())
}

func (e *tEstado) esValido(op int) bool {
	valido := false
	fa, ca := e.pos[0][a], e.pos[1][a]
	fb, cb := e.pos[0][b], e.pos[1][b]
	fc, cc := e.pos[0][c], e.pos[1][c]
	switch op {
	case arribaA:
		valido = fa > 1 && e.celdas[fa-2][ca] == vacia && e.celdas[fa-1][ca-1] == vacia && e.celdas[fa-1][ca+1] == vacia
		break
	case abajoA:
		valido = fa < dimension-2 && e.celdas[fa+2][ca] == vacia && e.celdas[fa+1][ca+1] == vacia && e.celdas[fa+1][ca-1] == vacia
		break
	case izquierdaA:
		valido = ca > 1 && e.celdas[fa][ca-2] == vacia && e.celdas[fa-1][ca-1] == vacia && e.celdas[fa+1][ca-1] == vacia
	case derechaA:
		valido = ca < dimension-2 && e.celdas[fa][ca+2] == vacia && e.celdas[fa+1][ca+1] == vacia && e.celdas[fa-1][ca+1] == vacia
		break
	case arribaB:
		valido = fb > 1 && e.celdas[fb-2][cb] == vacia && e.celdas[fb-1][cb-1] == vacia && e.celdas[fb-1][cb+1] == vacia
		break
	case abajoB:
		valido = fb < dimension-1 && e.celdas[fb+1][cb] == vacia && e.celdas[fb+1][cb-1] == vacia && e.celdas[fb+1][cb+1] == vacia
		break
	case izquierdaB:
		valido = cb > 1 && e.celdas[fb][cb-2] == vacia && e.celdas[fb-1][cb-1] == vacia
		break
	case derechaB:
		valido = cb < dimension-2 && e.celdas[fb][cb+2] == vacia && e.celdas[fb-1][cb+1] == vacia
		break
	case arribaC:
		valido = fc > 1 && e.celdas[fc-2][cc] == vacia
	case abajoC:
		valido = fc < dimension-2 && e.celdas[fc+2][cc] == vacia
	case izquierdaC:
		valido = cc > 0 && e.celdas[fc-1][cc-1] == vacia && e.celdas[fc][cc-1] == vacia && e.celdas[fc+1][cc-1] == vacia
		break
	case derechaC:
		valido = cc < dimension-1 && e.celdas[fc-1][cc+1] == vacia && e.celdas[fc][cc+1] == vacia && e.celdas[fc+1][cc+1] == vacia
		break
	}
	return valido
}

func (e *tEstado) aplicaOperador(op int) *tEstado {
	nuevo := *e
	switch op {
	case arribaA:
		nuevo.pos[0][a]--
		nuevo.moverAverticalmente(nuevo.pos[0][a], e.pos[0][a], e.pos[1][a], 1)
		break
	case abajoA:
		nuevo.pos[0][a]++
		nuevo.moverAverticalmente(nuevo.pos[0][a], e.pos[0][a], e.pos[1][a], -1)
		break
	case izquierdaA:
		nuevo.pos[1][a]--
		nuevo.moverAhorizontalmente(nuevo.pos[1][a], e.pos[1][a], e.pos[0][a], 1)
		break
	case derechaA:
		nuevo.pos[1][a]++
		nuevo.moverAhorizontalmente(nuevo.pos[1][a], e.pos[1][a], e.pos[0][a], -1)
		break
	case arribaB:
		nuevo.pos[0][b]--
		nuevo.moverBverticalmente(nuevo.pos[0][b], nuevo.pos[1][b], e.pos[0][b], 0)
		break
	case abajoB:
		nuevo.pos[0][b]++
		nuevo.moverBverticalmente(nuevo.pos[0][b], nuevo.pos[1][b], e.pos[0][b], 1)
		break
	case izquierdaB:
		nuevo.pos[1][b]--
		nuevo.moverBhorizontalmente(nuevo.pos[0][b], nuevo.pos[1][b], e.pos[1][b], 1)
		break
	case derechaB:
		nuevo.pos[1][b]++
		nuevo.moverBhorizontalmente(nuevo.pos[0][b], nuevo.pos[1][b], e.pos[1][b], -1)
		break
	case arribaC:
		nuevo.pos[0][c]--
		nuevo.moverCverticalmente(nuevo.pos[0][c], e.pos[0][c], nuevo.pos[1][c], 1)
		break
	case abajoC:
		nuevo.pos[0][c]++
		nuevo.moverCverticalmente(nuevo.pos[0][c], e.pos[0][c], nuevo.pos[1][c], -1)
	case izquierdaC:
		nuevo.pos[1][c]--
		nuevo.moverChorizontalmente(nuevo.pos[1][c], e.pos[1][c], nuevo.pos[0][c])
		break
	case derechaC:
		nuevo.pos[1][c]++
		nuevo.moverChorizontalmente(nuevo.pos[1][c], e.pos[1][c], nuevo.pos[0][c])
		break
	}
	return &nuevo
}

func (e *tEstado) moverAverticalmente(fn, fo, col, dir int) {
	e.celdas[fn][col+1] = a
	e.celdas[fn][col-1] = a
	e.celdas[fn-1*dir][col] = a
	e.celdas[fo+1*dir][col] = vacia
	e.celdas[fo][col+1] = vacia
	e.celdas[fo][col-1] = vacia
}

func (e *tEstado) moverAhorizontalmente(cn, co, fil, dir int) {
	e.celdas[fil-1][cn] = a
	e.celdas[fil+1][cn] = a
	e.celdas[fil][cn-1*dir] = a
	e.celdas[fil][co+1*dir] = vacia
	e.celdas[fil-1][co] = vacia
	e.celdas[fil+1][co] = vacia
}

func (e *tEstado) moverBverticalmente(fn, cn, fo, dir int) {
	e.celdas[fn-1+dir][cn] = b
	e.celdas[fn][cn+1] = b
	e.celdas[fn][cn-1] = b
	e.celdas[fo][cn+1] = vacia
	e.celdas[fo][cn-1] = vacia
	e.celdas[fo-dir][cn] = vacia
}

func (e *tEstado) moverBhorizontalmente(fn, cn, co, dir int) {
	e.celdas[fn-1][cn] = b
	e.celdas[fn][cn-1*dir] = b     // +1
	e.celdas[fn][co+1*dir] = vacia // -1
	e.celdas[fn-1][co] = vacia
}

func (e *tEstado) moverCverticalmente(fn, fo, col, dir int) {
	e.celdas[fn-1*dir][col] = c
	e.celdas[fo+1*dir][col] = vacia
}

func (e *tEstado) moverChorizontalmente(cn, co, fil int) {
	e.celdas[fil-1][cn] = c
	e.celdas[fil][cn] = c
	e.celdas[fil+1][cn] = c
	e.celdas[fil-1][co] = vacia
	e.celdas[fil][co] = vacia
	e.celdas[fil+1][co] = vacia
}

func (e *tEstado) coste(op int) int {
	return 1
}

func (e *tEstado) dispEstado() {
	fmt.Print("[")
	for i := 0; i < len(e.celdas)-1; i++ {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Println(e.celdas[i])
	}
	fmt.Print(" ", e.celdas[len(e.celdas)-1])
	fmt.Print("]\n")
}

func dispOperador(op int) {
	switch op {
	case arribaA:
		fmt.Println("ARRIBA_A:")
		break
	case abajoA:
		fmt.Println("ABAJO_A")
		break
	case izquierdaA:
		fmt.Println("IZQUIERDA_A:")
		break
	case derechaA:
		fmt.Println("DERECHA_A:")
		break
	case arribaB:
		fmt.Println("ARRIBA_B:")
		break
	case abajoB:
		fmt.Println("ABAJO_B:")
		break
	case izquierdaB:
		fmt.Println("IZQUIERDA_B:")
		break
	case derechaB:
		fmt.Println("DERECHA_B:")
		break
	case arribaC:
		fmt.Println("ARRIBA_C:")
		break
	case abajoC:
		fmt.Println("ABAJO_C:")
		break
	case izquierdaC:
		fmt.Println("IZQUIERDA_C:")
		break
	case derechaC:
		fmt.Println("DERECHA_C:")
		break
	}
}

func (e *tEstado) heuristica() int {
	return e.manhattan()
}

func (e *tEstado) manhattan() int {
	obj := estadoObjetivo()
	sum := 0.0
	for i := 1; i < len(e.pos[0]); i++ {
		sum += math.Abs(float64(e.pos[0][i]-obj.pos[0][i])) + math.Abs(float64(e.pos[1][i]-obj.pos[1][i]))
	}
	return int(sum)
}
