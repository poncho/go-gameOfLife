package main

import (
	"fmt"
)

type World struct {
	num int
	m   [20][20]string
}

func (w *World) initLife() {
	w.setCell(10, 11)
	w.setCell(10, 12)
	w.setCell(10, 13)
	w.setCell(11, 11)
	w.setCell(11, 12)

	w.fillWorld()
}

func (w World) printWorld() {
	for _, line := range w.m {
		for _, cell := range line {
			fmt.Print(cell)
		}

		fmt.Print("\n")
	}
}

func (w *World) fillWorld() {
	for i, _ := range w.m {
		for j, _ := range w.m[i] {
			if w.m[i][j] != "*" {
				w.m[i][j] = " "
			}
		}
	}
}

func (w *World) setCell(x int, y int) {
	w.m[x][y] = "*"
}

func (w *World) killCell(x int, y int) {
	w.m[x][y] = ""
}

func main() {
	w := &World{
		0,
		[20][20]string{},
	}

	w.initLife()
	w.printWorld()
}
