package snake

type Snake struct {
	Parent  *Snake
	Child   *Snake
	X       int
	Y       int
	Nodes   int
	Tick    int
	Growing bool
}

func (snake *Snake) Move(direction int) *Snake {
	returnPtr := snake.grow()
	current := snake
	for current.Parent != nil {
		current.X = current.Parent.X
		current.Y = current.Parent.Y
		current = current.Parent
	}
	current.setHeadCoordinate(direction)
	return returnPtr
}

func (snake *Snake) setHeadCoordinate(direction int) {
	switch direction {
	case 0:
		snake.Y--
	case 1:
		snake.Y++
	case 2:
		snake.X--
	case 3:
		snake.X++
	}
}

func (snake *Snake) grow() *Snake {
	if snake.Growing {
		snake.Tick--
	}
	if snake.Tick == 0 {
		tail := Snake{
			X: snake.X, Y: snake.Y, Parent: snake,
			Nodes: snake.Nodes + 1, Tick: snake.Nodes + 1, Growing: false,
		}
		snake.Child = &tail
		return &tail
	} else {
		return snake
	}
}
