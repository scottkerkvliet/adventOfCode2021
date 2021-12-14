package main

import (
	"Day13/origami"
	"fmt"
	"log"
)

func FirstFold(o *origami.Origami, folds []origami.Fold) {
	o.Fold(folds[0])

	fmt.Printf("After first fold, origami has %v points.\n", len(o.Points))
}

func main() {
	origami, folds, err := origami.ReadOrigami("origami.txt")
	if err != nil {
		log.Fatal(err)
	}

	FirstFold(origami, folds)
}
