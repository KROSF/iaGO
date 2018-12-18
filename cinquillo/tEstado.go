package main

const (
	min  = -1
	mesa = 0
	max  = 1
)

type tEstado struct {
	cartas [2][10]int
	max    int
	min    int
}

func estadoInicial(max []int, min []int) *tEstado {
	es := &tEstado{}
	for index, carta := range max {
		if carta != -1 {
			es.cartas[carta][index] = 1
			es.max++
		}
	}
	for index, carta := range min {
		if carta != -1 {
			es.cartas[carta][index] = -1
			es.min++
		}
	}
	return es
}
