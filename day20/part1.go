package main

import (
	"fmt"
	"log"
)

func getOuterPixels(algorithm string, current byte) byte {
	if current == '#' {
		return algorithm[511]
	}
	return algorithm[0]
}

func getIndexFromImage(image [][]byte, row, col int, outerPixels byte) int {
	index := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			index = index * 2
			if i >= 0 && i < len(image) && j >= 0 && j < len(image[i]) {
				if image[i][j] == '#' {
					index++
				}
			} else if outerPixels == '#' {
				index++
			}
		}
	}

	return index
}

func enhanceImage(algorithm string, image [][]byte, outerPixels byte) ([][]byte, byte) {
	var newImage [][]byte
	for i := -1; i < len(image)+1; i++ {
		newImage = append(newImage, make([]byte, len(image[0])+2))
		for j := -1; j < len(image[0])+1; j++ {
			algoIndex := getIndexFromImage(image, i, j, outerPixels)
			newImage[i+1][j+1] = algorithm[algoIndex]
		}
	}

	return newImage, getOuterPixels(algorithm, outerPixels)
}

func countLitPixels(algorithm string, image [][]byte) {
	cycles := 2
	outerPixels := byte('.')
	for i := 0; i < cycles; i++ {
		image, outerPixels = enhanceImage(algorithm, image, outerPixels)
	}

	pixelsLit := 0
	for _, row := range image {
		for _, pixel := range row {
			if pixel == '#' {
				pixelsLit++
			}
		}
	}

	fmt.Printf("After enhancing %v times, there were %v pixels lit.\n", cycles, pixelsLit)

	/*
		for _, row := range image {
			fmt.Println(string(row))
		}
	*/
}

func main() {
	algorithm, image, err := ReadImage("image.txt")
	if err != nil {
		log.Fatal(err)
	}

	countLitPixels(algorithm, image)
}
