package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var cellStr = "*"
var deadCellStr = " "

type World struct {
	num         int
	m           [20][20]string
	livingCells []Cell
	dyingCells  []Cell
}

type Cell struct {
	x int
	y int
}

func (w *World) initLife() {
	w.setCell(9, 12)
	w.setCell(10, 11)
	w.setCell(10, 12)
	w.setCell(10, 13)
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
			if w.m[i][j] != cellStr {
				w.m[i][j] = deadCellStr
			}
		}
	}
}

func (w World) printCells() {
	for i, _ := range w.m {
		for j, _ := range w.m[i] {
			if w.m[i][j] == cellStr {
				fmt.Println(i, j)
			}
		}
	}
}

func (w World) checkCellNeighbors(x int, y int) bool {
	neighbors := 0

	//LEFT
	if x-1 >= 0 && w.m[x-1][y] == cellStr {
		neighbors += 1
	}
	//RIGHT
	if x+1 < len(w.m) && w.m[x+1][y] == cellStr {
		neighbors += 1
	}
	//UP
	if y-1 >= 0 && w.m[x][y-1] == cellStr {
		neighbors += 1
	}
	//DOWN
	if y+1 < len(w.m[0]) && w.m[x][y+1] == cellStr {
		neighbors += 1
	}
	//UP LEFT
	if x-1 >= 0 && y-1 >= 0 && w.m[x-1][y-1] == cellStr {
		neighbors += 1
	}
	//UP RIGHT
	if x+1 < len(w.m) && y-1 >= 0 && w.m[x+1][y-1] == cellStr {
		neighbors += 1
	}
	//DOWN LEFT
	if x-1 >= 0 && y+1 < len(w.m[0]) && w.m[x-1][y+1] == cellStr {
		neighbors += 1
	}
	//DOWN RIGHT
	if x+1 < len(w.m) && y+1 < len(w.m[0]) && w.m[x+1][y+1] == cellStr {
		neighbors += 1
	}

	if w.m[x][y] == cellStr && (neighbors == 2 || neighbors == 3) {
		return true
	} else if w.m[x][y] == deadCellStr && neighbors == 3 {
		return true
	} else {
		return false
	}
}

func (w *World) setCell(x int, y int) {
	w.m[x][y] = cellStr
}

func (w *World) killCell(x int, y int) {
	w.m[x][y] = deadCellStr
}

func (w *World) startGame() {
	endGame := false
	duration := time.Second

	for endGame == false {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		w.printWorld()

		w.livingCells = []Cell{}
		w.dyingCells = []Cell{}
		for i, _ := range w.m {
			for j, _ := range w.m[i] {
				if w.checkCellNeighbors(i, j) {
					//fmt.Println(i, j, " LIVING")
					w.livingCells = append(w.livingCells, Cell{i, j})
				} else {
					//fmt.Println("DIE", i, j)
					w.dyingCells = append(w.dyingCells, Cell{i, j})
				}
			}
		}

		for _, cell := range w.livingCells {
			w.setCell(cell.x, cell.y)
		}

		for _, cell := range w.dyingCells {
			w.killCell(cell.x, cell.y)
		}

		time.Sleep(duration * 1)
	}
}

func main() {
	w := &World{
		0,
		[20][20]string{},
		make([]Cell, 0),
		make([]Cell, 0),
	}

	w.initLife()
	//w.printWorld()
	//w.printCells()
	w.startGame()
}
