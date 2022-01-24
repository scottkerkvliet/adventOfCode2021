package alu

type Value struct {
	reg   rune
	value int
}

func NewValue(val int) *Value {
	return &Value{0, val}
}

func NewRegValue(reg rune) *Value {
	return &Value{reg, 0}
}

func (v Value) IsRegValue() bool {
	return v.reg != 0
}

func (v Value) GetValue() int {
	return v.value
}

func (v Value) GetReg() rune {
	return v.reg
}
