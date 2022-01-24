package alu

import (
	"errors"
	"fmt"
)

type ALU struct {
	w, x, y, z int
}

func NewALU() *ALU {
	return &ALU{0, 0, 0, 0}
}

func (a *ALU) String() string {
	return fmt.Sprintf("w: %v,  x: %v,  y: %v,  z: %v", a.w, a.x, a.y, a.z)
}

func (a *ALU) RunProgram(program []*Instruction, inputs chan int) error {
	a.Reset()
	for i, instruction := range program {
		err := a.runInstruction(instruction, inputs)
		if err != nil {
			return fmt.Errorf("Instruction %v: %w", i, err)
		}
	}

	return nil
}

func (a *ALU) Reset() {
	a.w = 0
	a.x = 0
	a.y = 0
	a.z = 0
}

func (a *ALU) GetZ() int {
	return a.z
}

func (a *ALU) getValue(val *Value) int {
	if !val.IsRegValue() {
		return val.GetValue()
	}
	reg := a.getRegister(val)
	if reg == nil {
		return 0
	}

	return *reg
}

func (a *ALU) getRegister(val *Value) *int {
	if val == nil || !val.IsRegValue() {
		return nil
	}

	switch val.GetReg() {
	case 'w':
		return &(a.w)
	case 'x':
		return &(a.x)
	case 'y':
		return &(a.y)
	case 'z':
		return &(a.z)
	default:
		return nil
	}
}

func (a *ALU) runInstruction(instruction *Instruction, inputs chan int) error {
	switch instruction.Op {
	case Inp:
		return a.input(instruction.Val1, inputs)
	case Add:
		return a.add(instruction.Val1, instruction.Val2)
	case Mul:
		return a.multiply(instruction.Val1, instruction.Val2)
	case Div:
		return a.divide(instruction.Val1, instruction.Val2)
	case Mod:
		return a.modulo(instruction.Val1, instruction.Val2)
	case Eql:
		return a.equal(instruction.Val1, instruction.Val2)
	default:
		return nil
	}
}

func (a *ALU) input(val1 *Value, inputs chan int) error {
	reg := a.getRegister(val1)
	if reg == nil {
		return errors.New("Operator Input first operand is not a valid register.")
	}

	num, open := <-inputs
	if !open {
		return errors.New("Operator Input channel has already been closed.")
	}

	*reg = num
	return nil
}

func (a *ALU) add(val1, val2 *Value) error {
	reg := a.getRegister(val1)
	if reg == nil {
		return errors.New("Operator Add first operand is not a valid register.")
	}

	*reg = *reg + a.getValue(val2)
	return nil
}

func (a *ALU) multiply(val1, val2 *Value) error {
	reg := a.getRegister(val1)
	if reg == nil {
		return errors.New("Operator Multiply first operand is not a valid register.")
	}

	*reg = *reg * a.getValue(val2)
	return nil
}

func (a *ALU) divide(val1, val2 *Value) error {
	reg := a.getRegister(val1)
	if reg == nil {
		return errors.New("Operator Divide first operand is not a valid register.")
	}

	divisor := a.getValue(val2)
	if divisor == 0 {
		return errors.New("Operator Divide second operand is 0.")
	}

	*reg = *reg / divisor
	return nil
}

func (a *ALU) modulo(val1, val2 *Value) error {
	reg := a.getRegister(val1)
	if reg == nil {
		return errors.New("Operator Modulo first operand is not a valid register.")
	}

	divisor := a.getValue(val2)

	if divisor <= 0 {
		return fmt.Errorf("Operator Modulo second operand is less than or equal to 0 (%v).", divisor)
	}
	if *reg < 0 {
		return fmt.Errorf("Operator Modulo first operand is less than 0 (%v).", *reg)
	}

	*reg = *reg % divisor
	return nil
}

func (a *ALU) equal(val1, val2 *Value) error {
	reg := a.getRegister(val1)
	if reg == nil {
		return errors.New("Operator Equal first operand is not a valid register.")
	}

	if *reg == a.getValue(val2) {
		*reg = 1
	} else {
		*reg = 0
	}
	return nil
}
