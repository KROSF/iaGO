package main

import (
	"fmt"
)

type tNodo struct {
	estado      *TEstado
	padre       *tNodo
	operador    int
	coste       int
	profundidad int
	heuristica  int
}

type list = []*tNodo

func (n *tNodo) dispCamino() {
	if n.padre == nil {
		fmt.Printf("\nINICIO\n\n")
		n.estado.dispEstado()
	} else {
		n.padre.dispCamino()
		dispOperador(n.operador)
		n.estado.dispEstado()
	}
}

func (n *tNodo) dispSolucion() {
	n.dispCamino()
	fmt.Printf("Profundidad = %d\n", n.profundidad)
	fmt.Printf("Coste = %d\n", n.coste)
}

func (n *tNodo) dispNodo() {
	if n.padre != nil {
		fmt.Println("Padre:")
		n.padre.estado.dispEstado()
		fmt.Println("Estado:")
		n.estado.dispEstado()
		fmt.Println("Heuristica:", n.heuristica, "\nCoste:", n.coste)
		fmt.Println("Profundidad: ", n.profundidad)
		dispOperador(n.operador)
	}
}

func nodoInicial() *tNodo {
	return &tNodo{estadoInicial(), nil, 0, 0, 0, estadoInicial().heuristica()}
}

func nodoObjetivo() *tNodo {
	return &tNodo{estadoObjetivo(), nil, 0, 0, 0, 0}
}

func (n *tNodo) expadir() list {
	sucesores := list{}
	for op := 1; op <= operadores; op++ {
		if n.estado.esValido(op) {
			s := n.estado.aplicaOperador(op)
			sucesores = append(sucesores, &tNodo{s, n, op, n.coste + n.estado.coste(op), n.profundidad + 1, s.heuristica()})
		}
	}
	return sucesores
}

func (n *tNodo) existe(l list) bool {
	for i := 0; i < len(l); i++ {
		if l[i].estado.iguales(n.estado) {
			return true
		}
	}
	return false
}
