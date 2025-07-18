package main

type Rover struct {
	X         int
	Y         int
	Direction string // N, E, S, W
}

func (r *Rover) rotateLeft() {
	switch r.Direction {
	case "N":
		r.Direction = "W"
	case "E":
		r.Direction = "N"
	case "S":
		r.Direction = "E"
	case "W":
		r.Direction = "S"
	}
}

func (r *Rover) rotateRight() {
	switch r.Direction {
	case "N":
		r.Direction = "E"
	case "E":
		r.Direction = "S"
	case "S":
		r.Direction = "W"
	case "W":
		r.Direction = "N"
	}
}

func (r *Rover) moveForward(plateauWidth, plateauHeight int) {
	switch r.Direction {
	case "N":
		if r.Y < plateauHeight {
			r.Y++
		}
	case "E":
		if r.X < plateauWidth {
			r.X++
		}
	case "S":
		if r.Y > 0 {
			r.Y--
		}
	case "W":
		if r.X > 0 {
			r.X--
		}
	}
}

func (r *Rover) processInstructions(instructions string, plateauWidth, plateauHeight int) {
	for _, instruction := range instructions {
		switch instruction {
		case 'L':
			r.rotateLeft()
		case 'R':
			r.rotateRight()
		case 'M':
			r.moveForward(plateauWidth, plateauHeight)
		}
	}
}
