package main

import (
	"fmt"
)

const (
	torres     = 3
	aB         = 1
	aC         = 2
	bA         = 3
	bC         = 4
	cA         = 5
	cB         = 6
	a          = 0
	b          = 1
	c          = 2
	operadores = 6
)

var puzzleInicial = [torres][torres]int{
	{1, 2, 3},
	{0, 0, 0},
	{0, 0, 0},
}

var puzzleFinal = [torres][torres]int{
	{0, 0, 0},
	{0, 0, 0},
	{1, 2, 3},
}

type tEstado struct {
	towers [torres][torres]int
}

func crearEstado(puzzle [torres][torres]int) *tEstado {
	e := &tEstado{}
	e.towers = puzzle
	//fmt.Println(e.towers)
	return e
}

func estadoInicial() *tEstado {
	return crearEstado(puzzleInicial)
}

func estadoObjetivo() *tEstado {
	return crearEstado(puzzleFinal)
}

func (e *tEstado) iguales(f *tEstado) bool {
	for i := 0; i < torres; i++ {
		for j := 0; j < torres; j++ {
			if e.towers[i][j] != f.towers[i][j] {
				return false
			}
		}
	}
	return true
}

func (e *tEstado) testObjetivo() bool {
	return e.iguales(estadoObjetivo())
}

func (e *tEstado) piezasEnTorre(torre int) bool {
	piezas := 0
	for _, p := range e.towers[torre] {
		piezas += p
	}
	return piezas != 0
}

func (e *tEstado) esValido(op int) bool {
	valido := false
	switch op {
	case aB:
		valido = e.piezasEnTorre(a) && (e.towers[b][0] > e.towers[a][0] || e.towers[b][0] == 0)
		break
	case aC:
		valido = e.piezasEnTorre(a) && (e.towers[c][0] > e.towers[a][0] || e.towers[c][0] == 0)
		break
	case bA:
		valido = e.piezasEnTorre(b) && (e.towers[a][0] > e.towers[b][0] || e.towers[a][0] == 0)
		break
	case bC:
		valido = e.piezasEnTorre(b) && (e.towers[c][0] > e.towers[b][0] || e.towers[c][0] == 0)
		break
	case cA:
		valido = e.piezasEnTorre(c) && (e.towers[a][0] > e.towers[c][0] || e.towers[a][0] == 0)
		break
	case cB:
		valido = e.piezasEnTorre(c) && (e.towers[b][0] > e.towers[c][0] || e.towers[b][0] == 0)
		break
	}
	return valido
}

func (e *tEstado) moverFichas(nuevo *tEstado, from, to int) {
	if !e.piezasEnTorre(to) {
		nuevo.towers[to][0] = e.towers[from][0]
		for i := 0; i < torres-1; i++ {
			nuevo.towers[from][i] = nuevo.towers[from][i+1]
		}
		nuevo.towers[from][torres-1] = 0
	} else {
		for i := torres - 1; i > 0; i-- {
			nuevo.towers[to][i] = nuevo.towers[to][i-1]
		}
		nuevo.towers[to][0] = e.towers[from][0]
		for i := 0; i < torres-1; i++ {
			nuevo.towers[from][i] = nuevo.towers[from][i+1]
		}
		nuevo.towers[from][torres-1] = 0
	}
}

func (e *tEstado) aplicaOperador(op int) *tEstado {
	nuevo := *e
	switch op {
	case aB:
		e.moverFichas(&nuevo, a, b)
		break
	case aC:
		e.moverFichas(&nuevo, a, c)
		break
	case bA:
		e.moverFichas(&nuevo, b, a)
		break
	case bC:
		e.moverFichas(&nuevo, b, c)
		break
	case cA:
		e.moverFichas(&nuevo, c, a)
		break
	case cB:
		e.moverFichas(&nuevo, c, b)
		break
	}
	return &nuevo
}

func (e *tEstado) coste(op int) int {
	return 1
}

func (e *tEstado) dispEstado() {
	fmt.Print("[")
	for i := 0; i < torres-1; i++ {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Println(e.towers[i])
	}
	fmt.Print(" ", e.towers[torres-1])
	fmt.Print("]\n")
}

func dispOperador(op int) {
	switch op {
	case aB:
		fmt.Println("MOVER DE A a B")
		break
	case aC:
		fmt.Println("MOVER DE A a C")
		break
	case bA:
		fmt.Println("MOVER DE B a A")
		break
	case bC:
		fmt.Println("MOVER DE B a C")
		break
	case cA:
		fmt.Println("MOVER DE C a A")
		break
	case cB:
		fmt.Println("MOVER DE C a B")
		break
	}
}

func (e *tEstado) heuristica() int {
	piezas := 3
	for i := 0; i < torres; i++ {
		if e.towers[c][i] == i+1 {
			piezas--
		}
	}
	return piezas
}
