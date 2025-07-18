package main

import (
	"strings"
	"testing"
)

func TestRoverMovement(t *testing.T) {
	// Test 1: Simple movement North
	rover := Rover{X: 0, Y: 0, Direction: "N"}
	rover.moveForward(5, 5) // Moves to 0,1 N
	if rover.X != 0 || rover.Y != 1 || rover.Direction != "N" {
		t.Errorf("North movement failed: Expected 0 1 N, Got %d %d %s", rover.X, rover.Y, rover.Direction)
	}

	// Test 2: Rotate left (N -> W)
	rover.rotateLeft()
	if rover.Direction != "W" {
		t.Errorf("Left rotation failed: Expected W, Got %s", rover.Direction)
	}

	// Test 3: Rotate right (W -> N)
	rover.rotateRight()
	if rover.Direction != "N" {
		t.Errorf("Right rotation failed: Expected N, Got %s", rover.Direction)
	}

	// Test 4: Movement to the east and plateau boundary (5x5)
	roverEast := Rover{X: 4, Y: 0, Direction: "E"}
	roverEast.moveForward(5, 5) // Moves to 5,0 E
	if roverEast.X != 5 || roverEast.Y != 0 || roverEast.Direction != "E" {
		t.Errorf("East movement failed: Expected 5 0 E, Got %d %d %s", roverEast.X, roverEast.Y, roverEast.Direction)
	}
	roverEast.moveForward(5, 5) // Tries to move beyond the limit (should stay at 5,0 E)
	if roverEast.X != 5 || roverEast.Y != 0 || roverEast.Direction != "E" {
		t.Errorf("East boundary check failed: Expected 5 0 E, Got %d %d %s", roverEast.X, roverEast.Y, roverEast.Direction)
	}

	// Test 5: Movement to the south and plateau boundary (5x5)
	roverSouth := Rover{X: 0, Y: 1, Direction: "S"}
	roverSouth.moveForward(5, 5) // Moves to 0,0 S
	if roverSouth.X != 0 || roverSouth.Y != 0 || roverSouth.Direction != "S" {
		t.Errorf("South movement failed: Expected 0 0 S, Got %d %d %s", roverSouth.X, roverSouth.Y, roverSouth.Direction)
	}
	roverSouth.moveForward(5, 5) // Tries to move beyond the limit (should stay at 0,0 S)
	if roverSouth.X != 0 || roverSouth.Y != 0 || roverSouth.Direction != "S" {
		t.Errorf("South boundary check failed: Expected 0 0 S, Got %d %d %s", roverSouth.X, roverSouth.Y, roverSouth.Direction)
	}
}

func TestExecute(t *testing.T) {
	input := `5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
`
	expected := `1 3 N
5 1 E
`

	reader := strings.NewReader(input)
	var output strings.Builder

	Execute(reader, &output)

	if output.String() != expected {
		t.Errorf("Expected output:\n%q\nGot:\n%q", expected, output.String())
	}
}
