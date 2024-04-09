package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Board struct {
	grid          [][]bool
	width, height int
}

func main() {
	width, height := 60, 30

	currentBoard := generateRandomBoard(width, height)
	for {

		printBoard(currentBoard)

		var lastBoard Board = currentBoard

		currentBoard = generateNextBoard(lastBoard)

		time.Sleep(500 * time.Millisecond)
	}

}
func generateRandomBoard(width int, height int) Board {
	var board Board = generateEmptyBoard(width, height)
	for x := range board.width {
		for y := range board.height {
			if rand.Intn(20) == 0 {
				board.grid[y][x] = true
			}
		}
	}
	return board
}

/*
	Rules: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life

Any live cell with fewer than two live neighbors dies, as if by underpopulation.
Any live cell with two or three live neighbors lives on to the next generation.
Any live cell with more than three live neighbors dies, as if by overpopulation.
Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
*/
func generateNextBoard(currentBoard Board) Board {
	var newBoard Board = generateEmptyBoard(currentBoard.width, currentBoard.height)
	for x := range currentBoard.width {
		for y := range currentBoard.height {
			isAlive := currentBoard.grid[y][x]
			var aliveNeighbors int8 = countTrues(getNeighborValues(currentBoard, x, y))

			if isAlive && (aliveNeighbors == 2 || aliveNeighbors == 3) {
				setPixel(newBoard, x, y, true)
			} else {
				setPixel(newBoard, x, y, false)
			}
			if !isAlive && aliveNeighbors == 3 {
				setPixel(newBoard, x, y, true)
			}
		}
	}
	return newBoard
}
func countTrues(boolSlice []bool) int8 {
	var count int8 = 0
	for i := range boolSlice {
		if boolSlice[i] == true {
			count++
		}
	}
	return count
}

func getNeighborValues(board Board, x int, y int) []bool {
	neighborValues := []bool{}
	for _x := range 3 {
		_x -= 1
		for _y := range 3 {
			_y -= 1
			// check for out of bounds exceptions & not checking the passed pixel
			if (_x+x == x && _y+y == y) || (_x+x >= board.width || _y+y >= board.height) || (_x+x < 0 || _y+y < 0) {
				continue
			}
			neighborValues = append(neighborValues, board.grid[y+_y][x+_x])
		}
	}
	return neighborValues
}

// Cordinate grid (0,0) starts at top left
func setPixel(board Board, x int, y int, isAlive bool) {
	board.grid[y][x] = isAlive
}

func generateEmptyBoard(width int, height int) Board {
	newBoard := make([][]bool, height)

	for i := range height {
		newBoard[i] = make([]bool, width)
	}
	return Board{grid: newBoard, width: width, height: height}
}

func printBoard(board Board) {
	// This only works on Windows but I wear deodorant so it effecting Linux doesn't matter
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for row := range board.grid {
		for col := range board.grid[0] {
			var isPixelAlive bool = board.grid[row][col]
			if isPixelAlive {
				fmt.Print("#")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
}
