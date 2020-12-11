package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	input := readInput("input.txt")

	rows := convertInput(input)
	fmt.Println(len(rows))
	fmt.Println(len(rows[0]))
	createImage(rows)

	log.Println("completed in ", time.Since(start))
}

func createImage(input [][]string) {
	img := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	green := color.RGBA{0, 255, 0, 255}
	red := color.RGBA{255, 0, 0, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)
	for i, row := range input {
		for j, entry := range row {
			if entry == "1" {
				img.Set(i, j, red)
			}
		}
	}
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func convertInput(input []string) [][]string {

	var newRows [][]string
	for _, row := range input {
		row = strings.TrimPrefix(row, " [")
		row = strings.TrimSuffix(row, "],")
		rowSlice := strings.Split(row, ", ")
		// fmt.Printf("%+v (%T), %+v (%T)\n", rowSlice[0], rowSlice[0], rowSlice[999], rowSlice[999])
		newRows = append(newRows, rowSlice)
	}
	return newRows
}

// readInput will read the two wires routes and return as two strings
func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
