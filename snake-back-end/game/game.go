package game

import (
	"encoding/json"
	"math/rand"
	"snakeGame/snake"
	"time"

	"github.com/gorilla/websocket"
)

type Game struct {
	Direction    int `json:"-"`
	Grid         [][]int
	previousGrid [][]int
	Score        int
	GameOver     bool
	gameSpeed    time.Duration
	height       int
	width        int
	HasFood      bool
	foodX        int
	foodY        int
}

func NewGame(height int, width int) Game {
	grid := make([][]int, height)
	for y := range grid {
		grid[y] = make([]int, width)
		for x := range grid[y] {
			grid[y][x] = 0
		}
	}
	returnGrid := Game{Direction: 3, Grid: grid, Score: 0, gameSpeed: time.Second / 2, height: height, width: width, HasFood: false}
	returnGrid.addFood()
	return returnGrid
}

func (game *Game) Wipe() {
	time.Sleep(game.gameSpeed)
	for row := 0; row < game.height; row++ {
		for col := 0; col < game.width; col++ {
			if game.Grid[row][col] != 2 {
				game.Grid[row][col] = 0
			}
		}
	}
}

func (game *Game) DrawSnake(snakeTail *snake.Snake) {
	if game.SnakeIsAlive(snakeTail) {
		current := snakeTail
		for current.Parent != nil {
			game.Grid[current.Y][current.X] = 1
			current = current.Parent
		}
		game.Grid[current.Y][current.X] = 1
		snakeAte := current.X == game.foodX && current.Y == game.foodY
		if snakeAte {
			game.HasFood = false
			snakeTail.Growing = true
			game.addFood()
			game.addScore()
			game.gameSpeed -= time.Millisecond * 10
		}
	} else {
		game.setGameOver()
	}
}

func (game *Game) StorePreviousState() {
	game.previousGrid = make([][]int, len(game.Grid))
	for i := range game.Grid {
		game.previousGrid[i] = make([]int, len(game.Grid[i]))
		copy(game.previousGrid[i], game.Grid[i])
	}
}

func (game *Game) setGameOver() {
	game.GameOver = true
	game.Grid = game.previousGrid
	for row := 0; row < game.height; row++ {
		for col := 0; col < game.width; col++ {
			if game.Grid[row][col] != 1 {
				game.Grid[row][col] = -1
			}
		}
	}
}

func (game *Game) SnakeIsAlive(tail *snake.Snake) bool {
	head := tail.FindHead()
	if head.Y < 0 || head.Y >= len(game.Grid) || head.X < 0 || head.X >= len(game.Grid[0]) {
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

func (game *Game) addFood() {
	rand.Seed(time.Now().UnixNano())
	validPosition := false
	for !validPosition {
		randomRow := rand.Intn(len(game.Grid))
		randomCol := rand.Intn(len(game.Grid[0]))
		if game.Grid[randomRow][randomCol] == 0 {
			validPosition = true
			game.Grid[randomRow][randomCol] = 2
			game.foodX = randomCol
			game.foodY = randomRow
		}
	}
	game.HasFood = true
}

func (game *Game) addScore() {
	multiplier := (time.Second - game.gameSpeed) * 10
	game.Score += int(10 * multiplier.Seconds())
}

func (game *Game) Play(tail *snake.Snake, conn *websocket.Conn) {
	for !game.GameOver {
		game.StorePreviousState()
		game.Wipe()
		tail = tail.Move(game.Direction)
		game.Direction = -1
		game.DrawSnake(tail)
		jsonData, _ := json.Marshal(game)
		conn.WriteMessage(1, jsonData)
	}
}

func ComposeGameAndSnake() (*snake.Snake, Game) {
	//returns tail, grid for a game of snake
	tail := snake.Snake{X: 2, Y: 5, Nodes: 3, Tick: 3, Growing: false, Direction: 3}
	body := snake.Snake{X: 3, Y: 5, Nodes: 3, Tick: 3, Growing: false, Direction: 3}
	head := snake.Snake{X: 4, Y: 5, Nodes: 3, Tick: 3, Growing: false, Direction: 3}
	tail.Parent = &body
	body.Parent = &head
	gameInstance := NewGame(10, 10)
	return &tail, gameInstance
}
