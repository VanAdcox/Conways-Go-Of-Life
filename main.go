package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Board struct {
	grid          [][]bool
	width, height int
}

func main() {
	width, height := 60, 30

	currentBoard := generateEmptyBoard(width, height)
	// nextBoard := generateEmptyBoard(width, height)

	setPixel(currentBoard, 0, 0, true)
	setPixel(currentBoard, 1, 0, true)
	setPixel(currentBoard, 1, 1, true)
	setPixel(currentBoard, 2, 0, true)
	printBoard(currentBoard)
	fmt.Print(getNeighborValues(currentBoard, 1, 1))

}
func getNeighborValues(board Board, x int, y int) []bool {

	neighborValues := []bool{}

	for _x := range 3 {
		_x -= 1
		for _y := range 3 {
			_y -= 1

			// check for out of bounds exceptions & not checking the passed pixel
			if (_x+x == x && _y+y == y) || (_x+x > board.width || _y+y > board.height) || (_x+x < 0 || _y+y < 0) {
				continue
			}
			neighborValues = append(neighborValues, board.grid[x+_x][y+_y])
		}
		/*
			fmt.Printf("x: %v y: %v Alive: %t ", x, y+i, board.grid[x][y+i])
			fmt.Printf("x: %v y: %v Alive: %t ", x, y+i, board.gri[x][y+i])
		*/
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
