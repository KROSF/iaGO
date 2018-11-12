package main

import (
	"fmt"
)

const (
	ci         = 1
	cci        = 2
	mi         = 3
	mmi        = 4
	mci        = 5
	cd         = 6
	ccd        = 7
	md         = 8
	mmd        = 9
	mcd        = 10
	operadores = 10
)

var nomOP = []string{"CI", "CCI", "MI", "MMI", "CMI", "CD", "CCD", "MD", "MMD", "MCD"}

type misioneroCaninal struct {
	m, c int
}

func (m *misioneroCaninal) iguales(c *misioneroCaninal) bool {
	return m.m == c.m && m.c == c.c
}

type tEstado struct {
	izquierda, derecha misioneroCaninal
	barca              bool
}

func estadoInicial() *tEstado {
	return &tEstado{
		izquierda: misioneroCaninal{3, 3},
		derecha:   misioneroCaninal{0, 0},
		barca:     true,
	}
}

func estadoObjetivo() *tEstado {
	return &tEstado{
		izquierda: misioneroCaninal{0, 0},
		derecha:   misioneroCaninal{3, 3},
		barca:     false,
	}
}

func (e *tEstado) iguales(f *tEstado) bool {
	return e.derecha.iguales(&f.derecha) && e.izquierda.iguales(&f.izquierda) && (e.barca == f.barca)
}

func (e *tEstado) esObjetivo() bool {
	return e.iguales(estadoObjetivo())
}

func (e *tEstado) esValido(op int) bool {
	valido := false
	switch op {
	case ci, cci:
		valido = e.derecha.c > 0 && (e.izquierda.c+op <= e.izquierda.m || e.derecha.c-op <= e.derecha.m) && !e.barca
		break
	case mi, mmi:
		valido = e.derecha.m > 0 && (e.izquierda.c <= e.izquierda.m+(op-2) || e.derecha.c <= e.derecha.m-(op-2)) && !e.barca
		break
	case mci:
		valido = e.derecha.m > 0 && e.derecha.c > 0 && e.izquierda.c-1 <= e.izquierda.m-1 && !e.barca
		break
	case cd, ccd:
		valido = e.izquierda.c > 0 && (e.derecha.c+(op-5) <= e.derecha.m || e.izquierda.c-(op-5) <= e.izquierda.m) && e.barca
		break
	case md, mmd:
		valido = e.izquierda.m > 0 && (e.derecha.c <= e.derecha.m+(op-7) || e.izquierda.c <= e.izquierda.m-(op-7)) && e.barca
		break
	case mcd:
		valido = e.izquierda.m > 0 && e.izquierda.c > 0 && e.derecha.c+1 <= e.derecha.m+1 && e.barca
		break
	}
	return valido
}

func (e *tEstado) aplicaOperador(op int) *tEstado {
	nuevo := *e
	switch op {
	case ci:
		nuevo.izquierda.c++
		nuevo.derecha.c--
		nuevo.barca = true
		break
	case cci:
		nuevo.izquierda.c += 2
		nuevo.derecha.c -= 2
		nuevo.barca = true
	case mi:
		nuevo.izquierda.m++
		nuevo.derecha.m--
		nuevo.barca = true
		break
	case mmi:
		nuevo.izquierda.m += 2
		nuevo.derecha.m -= 2
		nuevo.barca = true
		break
	case mci:
		nuevo.izquierda.c++
		nuevo.derecha.c--
		nuevo.izquierda.m++
		nuevo.derecha.m--
		nuevo.barca = true
	case cd:
		nuevo.derecha.c++
		nuevo.izquierda.c--
		nuevo.barca = false
		break
	case ccd:
		nuevo.izquierda.c -= 2
		nuevo.derecha.c += 2
		nuevo.barca = false
		break
	case md:
		nuevo.izquierda.m--
		nuevo.derecha.m++
		nuevo.barca = false
		break
	case mmd:
		nuevo.izquierda.m -= 2
		nuevo.derecha.m += 2
		nuevo.barca = false
		break
	case mcd:
		nuevo.izquierda.c--
		nuevo.derecha.c++
		nuevo.izquierda.m--
		nuevo.derecha.m++
		nuevo.barca = false
		break
	}
	return &nuevo
}

func (e *tEstado) coste(op int) int {
	return 1
}

func (e *tEstado) dispEstado() {
	fmt.Printf("Izquierda: %+v\n", e.izquierda)
	fmt.Printf("Derecha: %+v\n", e.derecha)
	if e.barca {
		fmt.Println("Barca : Izquierda")
	} else {
		fmt.Println("Barca : Derecha")
	}
}

func dispOperador(op int) {
	fmt.Printf("%s\n\n", nomOP[op-1])
}

func (e *tEstado) heuristica() int {
	return (e.izquierda.m + e.izquierda.c) / 2
}
