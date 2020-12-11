package main

import (
	"bufio"
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

	part1(input)
	// part2(input)
	// part2alternate(input)

	log.Println("completed in ", time.Since(start))
}

func part1(input []string) {
	decodePassports(input)
}

func decodePassports(input []string) {
	ppCount := 0
	newPP := true

	var pp []string
	var pps [][]string

	for _, row := range input {
		if newPP == true {
			ppCount++
			newPP = false
		}
		if row == "" {
			pps = append(pps, pp)
			newPP = true
			pp = []string{}
			continue
		}
		//processpp
		entries := strings.Split(row, " ")
		// fmt.Println(entries)
		pp = append(pp, entries...)
		// fmt.Println(pp)

	}
	pps = append(pps, pp) // add the last one
	log.Println("Passport Count:", ppCount)

	validCount := 0
	for _, p := range pps {
		var byr, iyr, eyr, hgt, hcl, ecl, pid string
		// var byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
		for _, i := range p {
			j := strings.Split(i, ":")
			switch j[0] {
			case "byr":
				byr = j[1]
			case "iyr":
				iyr = j[1]
			case "eyr":
				eyr = j[1]
			case "hgt":
				hgt = j[1]
			case "hcl":
				hcl = j[1]
			case "ecl":
				ecl = j[1]
			case "pid":
				pid = j[1]
				// case "cid":
				// 	cid = j[1]
			}
		}
		// part 1
		// valid := byr != "" && iyr != "" && eyr != "" && hgt != "" && hcl != "" && ecl != "" && pid != "" // && cid != ""
		// part2
		valid := validDate(byr, 1920, 2002) && validDate(iyr, 2010, 2020) && validDate(eyr, 2020, 2030) && validHeight(hgt) && validHairColor(hcl) && validEyeColor(ecl) && validPID(pid) // && cid != ""
		if valid == true {
			validCount++
		}
	}
	log.Println("Valid:", validCount)
}

func validDate(year string, start, end int) bool {
	if year == "" {
		return false
	}
	yi, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	if yi < start || yi > end {
		return false
	}
	return true
}
func validHeight(hgt string) bool {
	if hgt == "" {
		return false
	}
	if !strings.HasSuffix(hgt, "cm") && !strings.HasSuffix(hgt, "in") {
		// log.Println("Fail Height", hgt)
		return false
	}
	if strings.HasSuffix(hgt, "cm") {
		val := strings.TrimSuffix(hgt, "cm")
		hi, err := strconv.Atoi(val)
		if err != nil {
			// log.Println("Fail Height", hgt)
			return false
		}
		if hi < 150 || hi > 193 {
			// log.Println("Fail Height", hgt)
			return false
		}
	}
	if strings.HasSuffix(hgt, "in") {
		val := strings.TrimSuffix(hgt, "in")
		hi, err := strconv.Atoi(val)
		if err != nil {
			// log.Println("Fail Height", hgt)
			return false
		}
		if hi < 59 || hi > 76 {
			// log.Println("Fail Height", hgt)
			return false
		}
	}
	return true
}
func validHairColor(hcl string) bool {
	if hcl == "" {
		return false
	}
	// if !strings.HasPrefix("hcl", "#") {
	// 	return false
	// }
	r := "^#([a-fA-F0-9]{6})$"
	match, _ := regexp.MatchString(r, hcl)
	if match == false {
		// log.Println("Failed HCL:", hcl)
		return false
	}
	return true
}
func validEyeColor(ecl string) bool {
	if ecl == "" {
		return false
	}
	validColors := "amb blu brn gry grn hzl oth"
	// definitely not the best way to check
	check := strings.Contains(validColors, ecl)
	if check == false {
		// log.Println("Failed Eye Color:", ecl)
		return false
	}
	// log.Println("Passed Eye Color:", ecl)
	return true
}
func validPID(pid string) bool {
	if pid == "" {
		return false
	}
	_, err := strconv.Atoi(pid)
	if err != nil {
		log.Println("Failed PID:", pid)
		return false
	}
	if len(pid) != 9 {
		log.Println("Failed PID:", pid)
		return false
	}
	return true
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
