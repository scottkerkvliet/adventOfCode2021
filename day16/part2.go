package main

import (
	fr "Day16/fileReader"
	pkt "Day16/packets"
	"fmt"
	"log"
)

func main() {
	transmission, err := fr.ReadTransmission("transmission.txt")
	if err != nil {
		log.Fatal(err)
	}

	packet, err := pkt.ReadPacket(transmission)
	if err != nil {
		log.Fatal(err)
	}

	value := packet.GetValue()
	fmt.Printf("The value of the transmission is %v.\n", value)
}
