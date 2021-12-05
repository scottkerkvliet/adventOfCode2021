package main

import (
	"fmt"
)

func main() {
	commands, err := ReadCommands("commands.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	horizontal, vertical, aim := 0, 0, 0

	for _, command := range commands {
		switch command.Direction {
		case Forward:
			horizontal += command.Length
			vertical += (aim * command.Length)
		case Down:
			aim += command.Length
		case Up:
			aim -= command.Length
		}
	}

	fmt.Printf("We have moved %v forward and have a depth of %v.\n", horizontal, vertical)
	fmt.Printf("Multiplied result is %v.\n", horizontal*vertical)
}
