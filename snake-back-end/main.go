package main

import (
	"log"
	"net/http"

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

var snakeGame game.Game
var tail *snake.Snake

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
		dir := int(message[0] - '0')
		if err == nil {
			snakeGame.Direction = dir
		}
	}
}

func snakeGameSocket(writer http.ResponseWriter, request *http.Request) {
	var upgrader = websocket.Upgrader{}
	var conn *websocket.Conn
	var err error
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	if conn, err = upgrader.Upgrade(writer, request, nil); err != nil {
		log.Fatal("failed to upgrade to websocket")
	}
	defer conn.Close()
	tail, snakeGame = game.ComposeGameAndSnake()
	snakeGame.Play(tail, conn)
}
func main() {
	http.HandleFunc("/snake", snakeGameSocket)
	http.HandleFunc("/input", snakeInputSocket)
	http.ListenAndServe(":8080", nil)
}
