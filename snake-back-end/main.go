package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"snakeGame/game"
	"snakeGame/snake"

	"github.com/gorilla/websocket"
)

const (
	Up = iota
	Down
	Left
	Right
)

var direction = 0

func newGame() (*snake.Snake, game.Grid) {
	//returns head, tail, grid for a game of snake
	tail := snake.Snake{X: 2, Y: 5, Nodes: 3, Tick: 3, Growing: false}
	body := snake.Snake{X: 3, Y: 5, Nodes: 3, Tick: 3, Growing: false} //, Child: &tail}
	head := snake.Snake{X: 4, Y: 5, Nodes: 3, Tick: 3, Growing: false} //, Child: &body}
	tail.Parent = &body
	body.Parent = &head
	grid := game.NewGrid(10, 10)
	return &tail, grid
}

func snakeInputSocket(writer http.ResponseWriter, request *http.Request) {
	var upgrader = websocket.Upgrader{}
	var conn *websocket.Conn
	var err error
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	if conn, err = upgrader.Upgrade(writer, request, nil); err != nil {
		log.Fatal("failed to upgrade to websocket (request new game endpoint)")
	}
	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("msg received error", err)
		}
		m := string(message)
		dir, err := strconv.Atoi(m)
		if err == nil {
			// direction = dir
			switch dir {
			case 0:
				if direction != Down {
					direction = dir
				}
			case 1:
				if direction != Up {
					direction = dir
				}
			case 2:
				if direction != Right {
					direction = dir
				}
			case 3:
				if direction != Left {
					direction = dir
				}
			}
		}
	}
}

func snakeGameSocket(writer http.ResponseWriter, request *http.Request) {
	var upgrader = websocket.Upgrader{}
	var conn *websocket.Conn
	var err error
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	if conn, err = upgrader.Upgrade(writer, request, nil); err != nil {
		log.Fatal("failed to upgrade to websocket (request new game endpoint)")
	}
	defer conn.Close()
	tail, grid := newGame()
	direction = Right
	for !grid.GameOver {
		grid.Wipe()
		directionThisCycle := direction
		tail = tail.Move(directionThisCycle)
		grid.DrawSnake(tail, directionThisCycle)
		jsonData, _ := json.Marshal(grid)
		conn.WriteMessage(1, jsonData)
	}
	direction = Right
}
func main() {
	http.HandleFunc("/snake", snakeGameSocket)
	http.HandleFunc("/input", snakeInputSocket)
	http.ListenAndServe(":8080", nil)
}
