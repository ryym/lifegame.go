package main

import (
	"fmt"
	"math/rand"
	"time"
)

type game struct {
	cells [][]bool
	nRows int
	nCols int
}

type pair struct {
	r int
	c int
}

func main() {
	game := newGame(30, 90)
	setRandomState(game)

	ticker := time.Tick(time.Second)
	for {
		<-ticker
		game.render()
		game.update()
	}
}

func newGame(nRows, nCols int, alives ...pair) *game {
	cells := make([][]bool, nRows)
	for r := 0; r < nRows; r++ {
		cells[r] = make([]bool, nCols)
	}

	for _, p := range alives {
		cells[p.r][p.c] = true
	}

	return &game{cells, nRows, nCols}
}

func (g *game) update() {
	for r, row := range g.cells {
		for c := 0; c < g.nCols; c++ {
			nAlives := countAliveCells(r, c, g)
			row[c] = computeNextState(g.cells[r][c], nAlives)
		}
	}
}

func (g *game) render() {
	fmt.Print("\033[0;0H")
	for r := 0; r < g.nRows; r++ {
		for c := 0; c < g.nCols; c++ {
			out := " "
			if g.cells[r][c] {
				out = "█"
			}
			fmt.Print(out)
		}
		fmt.Println()
	}
}

func countAliveCells(r, c int, g *game) int {
	neighbors := []pair{
		{r - 1, c - 1},
		{r - 1, c},
		{r - 1, c + 1},
		{r, c - 1},
		{r, c + 1},
		{r + 1, c - 1},
		{r + 1, c},
		{r + 1, c + 1},
	}

	nAlives := 0
	for _, p := range neighbors {
		if 0 <= p.r && p.r < g.nRows && 0 <= p.c && p.c < g.nCols {
			if g.cells[p.r][p.c] {
				nAlives++
			}
		}
	}
	return nAlives
}

func computeNextState(alive bool, nAlives int) bool {
	if alive {
		return nAlives == 2 || nAlives == 3
	} else {
		return nAlives == 3
	}
}

func setRandomState(g *game) {
	for r := 0; r < g.nRows; r++ {
		for c := 0; c < g.nCols; c++ {
			if rand.Intn(10) == 0 {
				g.cells[r][c] = true
			}
		}
	}
}
