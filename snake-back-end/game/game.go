package game

import (
	"math/rand"
	"snakeGame/snake"
	"time"
)

type Grid struct {
	Grid      [][]int
	Score     int
	GameOver  bool
	gameSpeed time.Duration
	height    int
	width     int
	HasFood   bool
	foodX     int
	foodY     int
}

func NewGrid(height int, width int) Grid {
	grid := make([][]int, height)
	for y := range grid {
		grid[y] = make([]int, width)
		for x := range grid[y] {
			grid[y][x] = 0
		}
	}
	returnGrid := Grid{Grid: grid, Score: 0, gameSpeed: time.Second / 2, height: height, width: width, HasFood: false}
	returnGrid.addFood()
	return returnGrid
}

func (grid *Grid) Wipe() {
	time.Sleep(grid.gameSpeed)
	for row := 0; row < grid.height; row++ {
		for col := 0; col < grid.width; col++ {
			if grid.Grid[row][col] != 2 {
				grid.Grid[row][col] = 0
			}
		}
	}
}

func (grid *Grid) DrawSnake(snakeTail *snake.Snake, direction int) {
	if grid.SnakeIsAlive(snakeTail) {
		current := snakeTail
		for current.Parent != nil {
			grid.Grid[current.Y][current.X] = 1
			current = current.Parent
		}
		grid.Grid[current.Y][current.X] = 1
		snakeAte := current.X == grid.foodX && current.Y == grid.foodY
		if snakeAte {
			grid.HasFood = false
			snakeTail.Growing = true
			grid.addFood()
			grid.addScore()
			grid.gameSpeed -= time.Millisecond * 10
		}
	} else {
		grid.setGameOver()
	}
}

func (grid *Grid) setGameOver() {
	grid.GameOver = true
	for row := 0; row < grid.height; row++ {
		for col := 0; col < grid.width; col++ {
			grid.Grid[row][col] = -1
		}
	}
}

func (grid *Grid) SnakeIsAlive(tail *snake.Snake) bool {
	head := tail.FindHead()
	println(head.X, head.Y)
	if head.Y < 0 || head.Y >= len(grid.Grid) || head.X < 0 || head.X >= len(grid.Grid[0]) {
		return false
	}
	current := tail
	for current.Parent != nil {
		if head.X == current.X && head.Y == current.Y {
			return false
		}
		current = current.Parent
	}
	return true
}

func (grid *Grid) addFood() {
	rand.Seed(time.Now().UnixNano())
	validPosition := false
	for !validPosition {
		randomRow := rand.Intn(len(grid.Grid))
		randomCol := rand.Intn(len(grid.Grid[0]))
		if grid.Grid[randomRow][randomCol] == 0 {
			validPosition = true
			grid.Grid[randomRow][randomCol] = 2
			grid.foodX = randomCol
			grid.foodY = randomRow
		}
	}
	grid.HasFood = true
}

func (grid *Grid) addScore() {
	multiplier := (time.Second - grid.gameSpeed) * 10
	grid.Score += int(10 * multiplier.Seconds())
}

func (grid *Grid) FoodPos() (int, int, bool) {
	return grid.foodX, grid.foodY, grid.HasFood
}
