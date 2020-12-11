package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	input := readInput("input.txt")

	result1 := part1(input)
	log.Println("Part1:", result1)

	result2 := part2(input)
	log.Println("Part2:", result2)

	log.Println("completed in ", time.Since(start))
}

func part1(input []string) int {
	bags := parseBags(input)
	bagsWithContainingBags := identifyContainingBags(bags)
	allBagsContainingBag := findBagsContainingBag("shiny gold", bagsWithContainingBags)
	fmt.Println(allBagsContainingBag)
	return len(allBagsContainingBag)
}

func part2(input []string) int {
	bags := parseBags(input)

	result := calculateNumberOfBagsRequired(bags, "shiny gold")

	return result
}

type Bag struct {
	Name     string
	Contains map[string]int
}

func calculateNumberOfBagsRequired(bags map[string]Bag, bagname string) int {
	count := 0
	for currentBagname, qty := range bags[bagname].Contains {
		fmt.Printf("adding %d to count for %s\n", qty, currentBagname)
		count += qty
		count += qty * calculateNumberOfBagsRequired(bags, currentBagname)
	}
	return count
}

func findBagsContainingBag(bagname string, bagsWithContainingBags map[string][]string) map[string]bool {
	bags := make(map[string]bool)
	if _, present := bagsWithContainingBags[bagname]; present {
		for _, containingBag := range bagsWithContainingBags[bagname] {
			bags[containingBag] = true
			for bag := range findBagsContainingBag(containingBag, bagsWithContainingBags) {
				bags[bag] = true
			}
		}
	}
	return bags
}

// identifyContainingBags takes the map of bagnames with containing bags and enumerates them to return a map
// of bag names and what they are contained in
func identifyContainingBags(bags map[string]Bag) map[string][]string {
	returnBags := make(map[string][]string)
	for containingBagname, bag := range bags {
		for containedBagName := range bag.Contains {
			returnBags[containedBagName] = append(returnBags[containedBagName], containingBagname)
		}
	}
	return returnBags
}

func parseBags(input []string) map[string]Bag {
	bags := make(map[string]Bag)
	for _, row := range input {
		bag := Bag{}
		bag.Contains = make(map[string]int)
		// First extract the top level bag
		re := regexp.MustCompile("^(.+) bags contain (.+)$")
		bagMatch := re.FindStringSubmatch(row)
		bag.Name = bagMatch[1]
		// Now extract all the contained bags
		bagsContained := strings.Split(bagMatch[2], ",")
		for _, containedBag := range bagsContained {
			if strings.Contains(containedBag, "no other") {
				continue
			}
			re2 := regexp.MustCompile("(\\d+) (.+) bag")
			bagMatch2 := re2.FindStringSubmatch(containedBag)
			containedBagName := bagMatch2[2]
			numBags, _ := strconv.Atoi(bagMatch2[1])
			bag.Contains[containedBagName] = numBags
		}
		// bags = append(bags, bag)
		bags[bag.Name] = bag
	}
	return bags
}

// readInput reads the input file and returns a slice of strings containing each line
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

// readRawInput returns a string of the entire input file
func readRawInput(path string) string {

	byteContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	content := string(byteContent)
	return content
}
