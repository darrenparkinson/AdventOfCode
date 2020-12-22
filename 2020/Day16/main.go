package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/darrenparkinson/AdventOfCode/pkg/aocutils"
)

func main() {

	ticket, tickets, fields, validNumbers := loadData()

	start := time.Now()
	result1 := part1(ticket, tickets, fields, validNumbers)
	log.Println("Result1:", result1)
	log.Println("completed in ", time.Since(start))

	start = time.Now()
	result2 := part2(ticket, tickets, fields, validNumbers)
	log.Println("Result2:", result2)
	log.Println("completed in ", time.Since(start))
}

func part1(myTicket Ticket, tickets []Ticket, fields []TicketField, validNumbers map[int]bool) int {

	_, errorRate := getValidTickets(tickets, validNumbers)
	return errorRate
}

func part2(myTicket Ticket, tickets []Ticket, fields []TicketField, validNumbers map[int]bool) int {
	validTickets, _ := getValidTickets(tickets, validNumbers)

	return len(validTickets)
}

func getValidTickets(tickets []Ticket, validNumbers map[int]bool) ([]Ticket, int) {
	ticketScanningErrorRate := 0
	var validTickets []Ticket
	for _, ticket := range tickets {
		ticketValid := true
		for _, value := range ticket.Values {
			if _, ok := validNumbers[value]; ok {
				continue
			}
			ticketValid = false
			ticketScanningErrorRate += value
			break
		}
		if ticketValid {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets, ticketScanningErrorRate
}

type TicketField struct {
	Name      string
	Range1Min int
	Range1Max int
	Range2Min int
	Range2Max int
}

type Ticket struct {
	Values []int
}

func loadData() (Ticket, []Ticket, []TicketField, map[int]bool) {
	myTicket := Ticket{aocutils.ReadInputRowsAsInts("input_myticket.txt")[0]}

	ticketrows := aocutils.ReadInputRowsAsInts("input_nearbytickets.txt")
	var tickets []Ticket
	for _, row := range ticketrows {
		ticket := Ticket{row}
		tickets = append(tickets, ticket)
	}

	fieldrows := aocutils.ReadInputAsRows("input_fields.txt")
	var fields []TicketField
	re := regexp.MustCompile(`^(.+): (.+) or (.+)$`)
	for _, row := range fieldrows {
		res := re.FindStringSubmatch(row)
		range1mins, range1maxs := strings.Split(res[2], "-")[0], strings.Split(res[2], "-")[1]
		range2mins, range2maxs := strings.Split(res[3], "-")[0], strings.Split(res[3], "-")[1]
		range1min, _ := strconv.Atoi(range1mins)
		range1max, _ := strconv.Atoi(range1maxs)
		range2min, _ := strconv.Atoi(range2mins)
		range2max, _ := strconv.Atoi(range2maxs)
		t := TicketField{res[1], range1min, range1max, range2min, range2max}
		fields = append(fields, t)
	}

	validNumbers := make(map[int]bool)
	for _, field := range fields {
		for i := field.Range1Min; i < field.Range1Max+1; i++ {
			validNumbers[i] = true
		}
		for j := field.Range2Min; j < field.Range2Max+1; j++ {
			validNumbers[j] = true
		}
	}

	return myTicket, tickets, fields, validNumbers
}
