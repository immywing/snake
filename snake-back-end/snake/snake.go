package snake

const (
	Up = iota
	Down
	Left
	Right
)

type Snake struct {
	Parent *Snake
	// Child   *Snake
	X         int
	Y         int
	Nodes     int
	Tick      int
	Growing   bool
	Direction int
}

func (snake *Snake) Move(direction int) *Snake {
	returnPtr := snake.grow()
	current := snake
	for current.Parent != nil {
		current.X = current.Parent.X
		current.Y = current.Parent.Y
		current = current.Parent
	}
	current.setDirection(direction)
	current.setHeadCoordinate()
	return returnPtr
}

func (snake *Snake) setDirection(direction int) {
	switch direction {
	case Up:
		if snake.Direction != Down {
			snake.Direction = 0
		}
	case Down:
		if snake.Direction != Up {
			snake.Direction = 1
		}
	case Left:
		if snake.Direction != Right {
			snake.Direction = 2
		}
	case Right:
		if snake.Direction != Left {
			snake.Direction = 3
		}
	}
}

func (snake *Snake) setHeadCoordinate() {
	switch snake.Direction {
	case Up:
		snake.Y--
	case Down:
		snake.Y++
	case Left:
		snake.X--
	case Right:
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
		return &tail
	} else {
		return snake
	}
}

func (snake *Snake) FindHead() *Snake {
	current := snake
	for current.Parent != nil {
		current = current.Parent
	}
	return current
}
