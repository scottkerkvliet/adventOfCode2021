package packets

import (
	tmn "Day16/transmission"
)

func BinaryToInt(binary string) (num int) {
	for _, bit := range binary {
		num = num * 2
		if bit == '1' {
			num++
		}
	}

	return
}

func ReadPacket(t *tmn.Transmission) (Packet, error) {
	versionString, err := t.ReadBits(3)
	if err != nil {
		return nil, err
	}
	typeString, err := t.ReadBits(3)
	if err != nil {
		return nil, err
	}
	version := BinaryToInt(versionString)
	typeId := BinaryToInt(typeString)

	var newPacket Packet
	if typeId == 4 {
		newPacket, err = readValuePacket(version, typeId, t)
	} else {
		newPacket, err = readOperatorPacket(version, typeId, t)
	}
	if err != nil {
		return nil, err
	}

	return newPacket, nil
}

func readValuePacket(version, typeId int, t *tmn.Transmission) (*ValuePacket, error) {
	end := false
	binaryValue := ""
	for !end {
		nextBits, err := t.ReadBits(5)
		if err != nil {
			return nil, err
		}
		end = nextBits[0] == '0'
		binaryValue += nextBits[1:]
	}

	return &ValuePacket{version, typeId, BinaryToInt(binaryValue)}, nil
}

func readOperatorPacket(version, typeId int, t *tmn.Transmission) (*OperatorPacket, error) {
	lengthType, err := t.ReadBits(1)
	if err != nil {
		return nil, err
	}
	var subPackets []Packet
	if lengthType == "0" {
		subPackets, err = readSubPacketsLength(t)
	} else {
		subPackets, err = readSubPacketsCount(t)
	}
	if err != nil {
		return nil, err
	}

	return &OperatorPacket{version, typeId, subPackets}, nil
}

func readSubPacketsLength(t *tmn.Transmission) ([]Packet, error) {
	lengthString, err := t.ReadBits(15)
	if err != nil {
		return nil, err
	}
	length := BinaryToInt(lengthString)
	subPacketString, err := t.ReadBits(length)
	if err != nil {
		return nil, err
	}

	subTransmission := tmn.NewBinaryTransmission(subPacketString)
	var subPackets []Packet
	var subPacket Packet
	for err == nil && !subTransmission.IsComplete() {
		subPacket, err = ReadPacket(subTransmission)
		if err == nil {
			subPackets = append(subPackets, subPacket)
		}
	}

	return subPackets, nil
}

func readSubPacketsCount(t *tmn.Transmission) ([]Packet, error) {
	countString, err := t.ReadBits(11)
	if err != nil {
		return nil, err
	}
	count := BinaryToInt(countString)

	var subPackets []Packet
	for i := 0; i < count; i++ {
		subPacket, err := ReadPacket(t)
		if err != nil {
			return nil, err
		}
		subPackets = append(subPackets, subPacket)
	}

	return subPackets, nil
}
