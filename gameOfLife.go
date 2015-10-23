package main

import (
	"fmt"
)

type World struct {
	num         int
	m           [20][20]string
	livingCells []Cell
}

type Cell struct {
	x int
	y int
}

func (w *World) initLife() {
	w.setCell(10, 11)
	w.setCell(10, 12)
	w.setCell(10, 13)

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

func (w World) printCells() {
	for _, cell := range w.livingCells {
		fmt.Println(cell)
	}
}

func (w *World) setCell(x int, y int) {
	w.m[x][y] = "*"
	w.livingCells = append(w.livingCells, Cell{x, y})
}

func (w *World) killCell(x int, y int) {
	w.m[x][y] = ""
}

func (w *World) startGame() {
	endGame := false

	for endGame == false {
		fmt.Println("TEST")
		endGame = true
	}
}

func main() {
	w := &World{
		0,
		[20][20]string{},
		make([]Cell, 0),
	}

	w.initLife()
	w.printWorld()
	w.printCells()
	w.startGame()
}
