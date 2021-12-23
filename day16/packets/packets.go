package packets

import "math"

type Packet interface {
	GetVersion() int
	GetTypeId() int
	GetValue() int
}

type ValuePacket struct {
	version, typeId, value int
}

func (vp ValuePacket) GetVersion() int {
	return vp.version
}

func (vp ValuePacket) GetTypeId() int {
	return vp.typeId
}

func (vp ValuePacket) GetValue() int {
	return vp.value
}

type OperatorPacket struct {
	version, typeId int
	subPackets      []Packet
}

func (op OperatorPacket) GetVersion() int {
	return op.version
}

func (op OperatorPacket) GetTypeId() int {
	return op.typeId
}

func (op OperatorPacket) GetValue() int {
	switch op.typeId {
	case 0:
		sum := 0
		for _, subPacket := range op.subPackets {
			sum += subPacket.GetValue()
		}
		return sum
	case 1:
		product := 1
		for _, subPacket := range op.subPackets {
			product = product * subPacket.GetValue()
		}
		return product
	case 2:
		min := math.MaxInt
		for _, subPacket := range op.subPackets {
			min = int(math.Min(float64(min), float64(subPacket.GetValue())))
		}
		return min
	case 3:
		max := math.MinInt
		for _, subPacket := range op.subPackets {
			max = int(math.Max(float64(max), float64(subPacket.GetValue())))
		}
		return max
	case 5:
		if op.subPackets[0].GetValue() > op.subPackets[1].GetValue() {
			return 1
		}
		return 0
	case 6:
		if op.subPackets[0].GetValue() < op.subPackets[1].GetValue() {
			return 1
		}
		return 0
	case 7:
		if op.subPackets[0].GetValue() == op.subPackets[1].GetValue() {
			return 1
		}
		return 0
	}

	return 0
}

func (op OperatorPacket) GetSubPackets() []Packet {
	return op.subPackets
}
