package main

import (
	"fmt"
)

var cellStr = "*"

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
			if w.m[i][j] != cellStr {
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

func (w World) checkCellNeighbors(x int, y int) bool {
	neighbors := 0

	//LEFT
	if x-1 >= 0 && w.m[x-1][y] == cellStr {
		neighbors += 1
	}
	//RIGHT
	if x+1 < len(w.m[0]) && w.m[x-1][y] == cellStr {
		neighbors += 1
	}
	//UP
	if y-1 >= 0 && w.m[x][y-1] == cellStr {
		neighbors += 1
	}
	//DOWN
	if y+1 >= len(w.m[0][0]) && w.m[x][y+1] == cellStr {
		neighbors += 1
	}
	//UP LEFT
	if x-1 >= 0 && y-1 >= 0 && w.m[x-1][y-1] == cellStr {
		neighbors += 1
	}
	//UP RIGHT
	if x+1 < len(w.m[0]) && y-1 >= 0 && w.m[x+1][y-1] == cellStr {
		neighbors += 1
	}
	//DOWN LEFT
	if x-1 >= 0 && y+1 < len(w.m[0][0]) && w.m[x-1][y+1] == cellStr {
		neighbors += 1
	}
	//DOWN RIGHT
	if x+1 >= len(w.m[0]) && y+1 >= len(w.m[0][0]) && w.m[x-1][y-1] == cellStr {
		neighbors += 1
	}

	if neighbors == 3 {
		return true
	} else {
		return false
	}
}

func (w *World) setCell(x int, y int) {
	w.m[x][y] = cellStr
	w.livingCells = append(w.livingCells, Cell{x, y})
}

func (w *World) killCell(x int, y int) {
	w.m[x][y] = ""
}

func (w *World) startGame() {
	endGame := false
	//var quit string
	//var i int

	for endGame == false {
		w.printCells()
		for i, _ := range w.m {
			for j, _ := range w.m[i] {
				if w.m[i][j] == cellStr {
					if w.checkCellNeighbors(i, j) {

					}
				} else {
					fmt.Println("DEAD")
				}
			}
		}
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
	//w.printWorld()
	//w.printCells()
	w.startGame()
}
