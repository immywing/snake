package game_test

import (
	"snakeGame/game"
	"snakeGame/snake"
	"testing"
)

func TestSnakeIsAliveWithValidPositions(t *testing.T) {
	tail := snake.Snake{X: 0, Y: 0, Nodes: 3, Tick: 3, Growing: false}
	body := snake.Snake{X: 1, Y: 0, Nodes: 3, Tick: 3, Growing: false} //, Child: &tail}
	head := snake.Snake{X: 2, Y: 0, Nodes: 3, Tick: 3, Growing: false} //, Child: &body}
	tail.Parent = &body
	body.Parent = &head
	grid := game.NewGrid(10, 10)
	got := grid.SnakeIsAlive(&tail)

	if !got {
		t.Errorf("Expected true, got: %v", got)
	}
}

func TestSnakeIsAliveWhenOutOfArrayIndex(t *testing.T) {
	tail := snake.Snake{X: 0, Y: 0, Nodes: 3, Tick: 3, Growing: false}
	body := snake.Snake{X: 1, Y: 0, Nodes: 3, Tick: 3, Growing: false}  //, Child: &tail}
	head := snake.Snake{X: 1, Y: -1, Nodes: 3, Tick: 3, Growing: false} //, Child: &body}
	tail.Parent = &body
	body.Parent = &head
	grid := game.NewGrid(10, 10)
	got := grid.SnakeIsAlive(&tail)

	if got {
		t.Errorf("Expected false, got: %v", got)
	}
}

func TestSnakeIsAliveWhenPtrsHaveSameCoordinates(t *testing.T) {
	tail := snake.Snake{X: 0, Y: 0, Nodes: 3, Tick: 3, Growing: false}
	body := snake.Snake{X: 1, Y: 0, Nodes: 3, Tick: 3, Growing: false} //, Child: &tail}
	head := snake.Snake{X: 1, Y: 0, Nodes: 3, Tick: 3, Growing: false} //, Child: &body}
	tail.Parent = &body
	body.Parent = &head
	grid := game.NewGrid(10, 10)
	got := grid.SnakeIsAlive(&tail)

	if got {
		t.Errorf("Expected false, got: %v", got)
	}
}

func TestSnakeIsAliveWhenPtrsHaveSameCoordinatesAndOutOfArray(t *testing.T) {
	tail := snake.Snake{X: 0, Y: 0, Nodes: 3, Tick: 3, Growing: false}
	body := snake.Snake{X: -1, Y: 0, Nodes: 3, Tick: 3, Growing: false} //, Child: &tail}
	head := snake.Snake{X: -1, Y: 0, Nodes: 3, Tick: 3, Growing: false} //, Child: &body}
	tail.Parent = &body
	body.Parent = &head
	grid := game.NewGrid(10, 10)
	got := grid.SnakeIsAlive(&tail)

	if got {
		t.Errorf("Expected false, got: %v", got)
	}
}
