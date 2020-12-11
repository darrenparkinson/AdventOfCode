package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i := 1; i < 25; i++ {
		day := strconv.Itoa(i)
		if len(day) < 2 {
			day = fmt.Sprintf("0%s", day)
		}
		path := fmt.Sprintf("Day%s", day)
		puzzleA := fmt.Sprintf("%s/a", path)
		puzzleB := fmt.Sprintf("%s/b", path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, os.ModeDir|0755)
			os.Mkdir(puzzleA, os.ModeDir|0755)
			os.Mkdir(puzzleB, os.ModeDir|0755)
		}
	}
}
