package main

import (
	"fmt"
	"log"
)

func sumFish(schedule *[9]int) (sum int) {
	for _, fishCount := range *schedule {
		sum += fishCount
	}
	return
}

func runSchedule(schedule *[9]int, days int) {
	for i := 0; i < days; i++ {
		births := schedule[0]
		for j := 1; j < len(*schedule); j++ {
			schedule[j-1] = schedule[j]
		}
		schedule[8] = births
		schedule[6] += births
	}
}

func getFishSchedule(fishSlice *[]int) *[9]int {
	schedule := [9]int{}
	for _, fish := range *fishSlice {
		schedule[fish] += 1
	}
	return &schedule
}

func main() {
	fish, err := ReadFish("fish.txt")
	if err != nil {
		log.Fatal(err)
	}

	days := 80
	schedule := getFishSchedule(fish)
	runSchedule(schedule, days)
	fmt.Printf("After %v days, there are %v fish\n", days, sumFish(schedule))
}
