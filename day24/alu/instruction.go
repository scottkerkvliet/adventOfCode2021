package alu

import (
	"strconv"
)

type Operation int

const (
	Nil Operation = iota
	Inp
	Add
	Mul
	Div
	Mod
	Eql
)

type Instruction struct {
	Op         Operation
	Val1, Val2 *Value
}

func NewInstruction(op string, val1, val2 string) *Instruction {
	operation := determineOperation(op)
	value1 := createValue(val1)
	value2 := createValue(val2)

	return &Instruction{operation, value1, value2}
}

func determineOperation(op string) Operation {
	switch op {
	case "inp":
		return Inp
	case "add":
		return Add
	case "mul":
		return Mul
	case "div":
		return Div
	case "mod":
		return Mod
	case "eql":
		return Eql
	default:
		return Nil
	}
}

func createValue(val string) *Value {
	num, err := strconv.Atoi(val)
	if err == nil {
		return NewValue(num)
	}

	if len(val) == 1 {
		return NewRegValue([]rune(val)[0])
	}

	return NewValue(0)
}
