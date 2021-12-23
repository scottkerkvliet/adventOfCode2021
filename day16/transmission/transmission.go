package transmission

import (
	"errors"
)

var hexLookup = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func hexToBinary(hex string) (binary string) {
	for _, hexChar := range hex {
		binary += hexLookup[hexChar]
	}
	return
}

type Transmission struct {
	binaryBuffer string
}

func NewHexTransmission(buffer string) *Transmission {
	return &Transmission{hexToBinary(buffer)}
}

func NewBinaryTransmission(buffer string) *Transmission {
	return &Transmission{buffer}
}

func (t *Transmission) ReadBits(num int) (string, error) {
	if len(t.binaryBuffer) < num {
		return "", errors.New("Not enough bits remaining in transmission")
	}

	bits := t.binaryBuffer[:num]
	t.binaryBuffer = t.binaryBuffer[num:]

	return bits, nil
}

func (t *Transmission) IsComplete() bool {
	return len(t.binaryBuffer) == 0
}
