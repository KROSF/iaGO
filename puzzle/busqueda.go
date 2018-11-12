package main

import (
	"sort"
)

type evaluacion func(list, int, int) bool

func vorax(l list, i, j int) bool {
	return l[i].heuristica < l[j].heuristica
}

func aestrella(l list, i, j int) bool {
	return l[i].heuristica+l[i].coste < l[j].heuristica+l[j].coste
}

func busqueda() {
	objetivo := false
	inicial := nodoInicial()
	actual := &tNodo{}
	abiertos := list{}
	abiertos = append(abiertos, inicial)
	for len(abiertos) != 0 && !objetivo {
		actual, abiertos = abiertos[0], abiertos[1:]
		objetivo = actual.estado.testObjetivo()
		if !objetivo {
			sucesores := actual.expadir()
			abiertos = append(abiertos, sucesores...)
		}
	}
	actual.dispSolucion()
}

func busquedaAnchura() {
	objetivo := false
	inicial := nodoInicial()
	actual := &tNodo{}
	abiertos := list{}
	cerrados := list{}
	abiertos = append(abiertos, inicial)
	for len(abiertos) != 0 && !objetivo {
		actual, abiertos = abiertos[0], abiertos[1:]
		objetivo = actual.estado.testObjetivo()
		repetido := actual.existe(cerrados)
		actual.dispNodo()
		if !objetivo && !repetido {
			sucesores := actual.expadir()
			abiertos = append(abiertos, sucesores...)
		}
		cerrados = append(cerrados, actual)
	}
	actual.dispSolucion()
}

func busquedaProfundidad() {
	objetivo := false
	inicial := nodoInicial()
	actual := &tNodo{}
	abiertos := list{}
	cerrados := list{}
	abiertos = append(abiertos, inicial)
	for len(abiertos) != 0 && !objetivo {
		actual, abiertos = abiertos[0], abiertos[1:]
		objetivo = actual.estado.testObjetivo()
		repetido := actual.existe(cerrados)
		if !objetivo && !repetido {
			sucesores := actual.expadir()
			abiertos = append(sucesores, abiertos...)
		}
		cerrados = append(cerrados, actual)
	}
	actual.dispSolucion()
}

func busquedaInformada(cmp evaluacion) {
	objetivo := false
	inicial := nodoInicial()
	actual := &tNodo{}
	abiertos := list{}
	cerrados := list{}
	abiertos = append(abiertos, inicial)
	for len(abiertos) != 0 && !objetivo {
		actual, abiertos = abiertos[0], abiertos[1:]
		objetivo = actual.estado.testObjetivo()
		repetido := actual.existe(cerrados)
		if !objetivo && !repetido {
			sucesores := actual.expadir()
			abiertos = append(sucesores, abiertos...)
			sort.Slice(abiertos, func(i, j int) bool {
				return cmp(abiertos, i, j)
			})
			cerrados = append(cerrados, actual)
		}
	}
	actual.dispSolucion()
}
