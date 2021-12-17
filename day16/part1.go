package main

import (
	fr "Day16/fileReader"
	pkt "Day16/packets"
	"fmt"
	"log"
)

func sumVersionNumbers(p pkt.Packet) int {
	sum := p.GetVersion()
	op, isType := p.(*pkt.OperatorPacket)
	if isType {
		for _, subPacket := range op.GetSubPackets() {
			sum += sumVersionNumbers(subPacket)
		}
	}

	return sum
}

func main() {
	transmission, err := fr.ReadTransmission("transmission.txt")
	if err != nil {
		log.Fatal(err)
	}

	packet, err := pkt.ReadPacket(transmission)
	if err != nil {
		log.Fatal(err)
	}

	sum := sumVersionNumbers(packet)
	fmt.Printf("The sum of the version numbers is %v.\n", sum)
}
