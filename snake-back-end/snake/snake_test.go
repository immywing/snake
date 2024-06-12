package snake_test

import (
	"reflect"
	"snakeGame/snake"
	"testing"
)

func TestMoveUpOnce(t *testing.T) {
	tail := snake.Snake{X: 0, Y: 0}
	body := snake.Snake{X: 1, Y: 0}
	head := snake.Snake{X: 2, Y: 0}
	tail.Parent = &body
	body.Parent = &head
	tail.Move(0)
	wantX := []int{1, 2, 2}
	wantY := []int{0, 0, -1}
	gotX := []int{tail.X, body.X, head.X}
	gotY := []int{tail.Y, body.Y, head.Y}
	if !reflect.DeepEqual(wantX, gotX) || !reflect.DeepEqual(wantY, gotY) {
		t.Errorf("Expected X positions: %v, got: %v", wantX, gotX)
		t.Errorf("Expected Y positions: %v, got: %v", wantY, gotY)
	}
}

func TestMoveDownOnce(t *testing.T) {
	tail := snake.Snake{X: 0, Y: 0}
	body := snake.Snake{X: 1, Y: 0}
	head := snake.Snake{X: 2, Y: 0}
	tail.Parent = &body
	body.Parent = &head
	tail.Move(1)
	wantX := []int{1, 2, 2}
	wantY := []int{0, 0, 1}
	gotX := []int{tail.X, body.X, head.X}
	gotY := []int{tail.Y, body.Y, head.Y}
	if !reflect.DeepEqual(wantX, gotX) || !reflect.DeepEqual(wantY, gotY) {
		t.Errorf("Expected X positions: %v, got: %v", wantX, gotX)
		t.Errorf("Expected Y positions: %v, got: %v", wantY, gotY)
	}
}

func TestMoveLeftOnce(t *testing.T) {
	tail := snake.Snake{X: 2, Y: 0}
	body := snake.Snake{X: 1, Y: 0}
	head := snake.Snake{X: 0, Y: 0}
	tail.Parent = &body
	body.Parent = &head
	tail.Move(2)
	wantX := []int{1, 0, -1}
	wantY := []int{0, 0, 0}
	gotX := []int{tail.X, body.X, head.X}
	gotY := []int{tail.Y, body.Y, head.Y}
	if !reflect.DeepEqual(wantX, gotX) || !reflect.DeepEqual(wantY, gotY) {
		t.Errorf("Expected X positions: %v, got: %v", wantX, gotX)
		t.Errorf("Expected Y positions: %v, got: %v", wantY, gotY)
	}
}

func TestMoveRightOnce(t *testing.T) {
	tail := snake.Snake{X: 0, Y: 0}
	body := snake.Snake{X: 1, Y: 0}
	head := snake.Snake{X: 2, Y: 0}
	tail.Parent = &body
	body.Parent = &head
	tail.Move(3)
	wantX := []int{1, 2, 3}
	wantY := []int{0, 0, 0}
	gotX := []int{tail.X, body.X, head.X}
	gotY := []int{tail.Y, body.Y, head.Y}
	if !reflect.DeepEqual(wantX, gotX) || !reflect.DeepEqual(wantY, gotY) {
		t.Errorf("Expected X positions: %v, got: %v", wantX, gotX)
		t.Errorf("Expected Y positions: %v, got: %v", wantY, gotY)
	}
}
