package main

import (
	"fmt"
	"io"
	"strings"
)

func Execute(inputReader io.Reader, outputWriter io.Writer) {
	var plateauWidth, plateauHeight int
	fmt.Fscan(inputReader, &plateauWidth, &plateauHeight)

	for {
		var startX, startY int
		var startDir string

		_, err := fmt.Fscan(inputReader, &startX, &startY, &startDir)
		if err != nil {
			break
		}

		var instructions string
		fmt.Fscan(inputReader, &instructions)

		rover := Rover{
			X:         startX,
			Y:         startY,
			Direction: startDir,
		}

		rover.processInstructions(instructions, plateauWidth, plateauHeight)
		fmt.Fprintf(outputWriter, "%d %d %s\n", rover.X, rover.Y, rover.Direction)
	}
}

func main() {

	inputReader := strings.NewReader(MockInputString)

	var capturedOutput strings.Builder
	outputWriter := &capturedOutput

	Execute(inputReader, outputWriter)

	fmt.Print(capturedOutput.String())
}
