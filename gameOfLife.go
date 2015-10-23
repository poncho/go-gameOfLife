package main

import (
	"fmt"
	"os"
	"os/exec"
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

func (w *World) checkCell(c Cell) int {
	livingNeighbors := 0

	if c.x-1 >= 0 {
		if w.m[c.x-1][c.y] == "*" {
			livingNeighbors += 1
		}
	}

	if c.x+1 <= len(w.m)-1 {
		if w.m[c.x+1][c.y] == "*" {
			livingNeighbors += 1
		}
	}

	if c.y-1 >= 0 {
		if w.m[c.x][c.y-1] == "*" {
			livingNeighbors += 1
		}
	}

	if c.y+1 <= len(w.m)-1 {
		if w.m[c.x][c.y+1] == "*" {
			livingNeighbors += 1
		}
	}
	fmt.Println(c.x, ",", c.y, " : ", livingNeighbors)
	return livingNeighbors
}

func (w *World) checkDeadCells() {

	for i, _ := range w.m {
		for j, _ := range w.m[i] {
			if w.m[i][j] == " " {
				if w.checkCell(Cell{i, j}) == 3 {
					//w.setCell(i, j)
				}
			}
		}
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
	var quit string
	//var i int

	for endGame == false {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		w.printWorld()

		w.checkDeadCells()

		/*
			for _, cell := range w.livingCells {
				_ = w.checkCell(cell)
			}
		*/

		fmt.Scanf("%s", &quit)

		if quit == "q" || quit == "Q" {
			endGame = true
		}
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
