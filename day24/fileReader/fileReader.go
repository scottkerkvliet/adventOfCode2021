package filereader

import (
	"Day24/alu"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInstructions(f string) ([]*alu.Instruction, error) {
	instructionFile, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %v", f)
	}

	var instructions []*alu.Instruction
	scanner := bufio.NewScanner(instructionFile)
	for scanner.Scan() {
		instructionLine := scanner.Text()
		instructionArgs := strings.Split(instructionLine, " ")
		if len(instructionArgs) < 2 || len(instructionArgs) > 3 {
			return nil, fmt.Errorf("Invalid number of arguments: \"%v\"", instructionLine)
		}
		var thirdArg string
		if len(instructionArgs) == 3 {
			thirdArg = instructionArgs[2]
		}

		instruction := alu.NewInstruction(instructionArgs[0], instructionArgs[1], thirdArg)
		instructions = append(instructions, instruction)
	}

	return instructions, nil
}
