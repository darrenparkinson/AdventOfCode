package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {

	// input := aocutils.ReadInputAsRows("input.txt")

	// start := time.Now()
	// result1 := part1(input)
	// log.Println("Result1:", result1)
	// log.Println("completed in ", time.Since(start))

	// start = time.Now()
	// result2 := part2(input)
	// log.Println("Result2:", result2)
	// log.Println("completed in ", time.Since(start))

	// test input should be 26, 437, 12240, 13632 respectively
	input := []string{"2 * 3 + (4 * 5)", "5 + (8 * 3 + 9 + 3 * 4 * 3)", "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"}
	fmt.Println(part1(input))
}
func part1(input []string) int {

	for _, row := range input {
		rpn := getRPN(row)
		fmt.Println(row, rpn)
	}
	rpn := getRPN(input[0])
	calc(rpn)

	return len(input)
}

func part2(input []string) int {
	return len(input)
}

// given a stack of strings in RPN format calculate the result
func calc(input StringStack) {
	// var stack StringStack
	for _, token := range input {
		if token == "*" || token == "+" {
			b, _ := input.Pop()
			op2 := atoi(b)
			a, _ := input.Pop()
			op1 := atoi(a)
			var result int
			switch token {
			case "+":
				result = op1 + op2
			case "*":
				result = op1 * op2
			}
			fmt.Println(result)
		}
	}
}

// atoi to just give back the int -- so bored of having to handle the error from strconv lol
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err) // we must have made a mistake
	}
	return i
}

// getRPN takes a string of tokens and parses them using the shunting-yard method to obtain reverse polish notation
// https://en.wikipedia.org/wiki/Shunting-yard_algorithm
func getRPN(tokens string) StringStack {
	fmt.Println(tokens)
	tokens = reverse(strings.ReplaceAll(tokens, " ", ""))
	tokens = strings.ReplaceAll(tokens, " ", "")

	// First push everything onto stacks
	var stack StringStack
	var output StringStack
	for _, r := range tokens {
		switch val := string(r); val {
		case "*", "+":
			stack.Push(val)
		case ")":
			stack.Push(val)
		case "(":
			i, _ := stack.Pop()
			for i != ")" {
				output.Push(i)
				i, _ = stack.Pop()
			}
			// j, _ := stack.Pop() // discard paren
			// fmt.Println("Discarding", j)
		default:
			output.Push(val)
		}
	}
	// push remaining stack to output -- this is our RPN (without precedence)
	for stack.IsEmpty() == false {
		i, _ := stack.Pop()
		output.Push(i)
	}
	// fmt.Println(tokens, stack, output, len(stack), len(output))
	return output
}

func reverse(str string) string {
	var result string
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

// StringStack represents a stack structure for LIFO with Push, Pop and IsEmpty methods
type StringStack []string

// IsEmpty checks to see if the stack is empty
func (s *StringStack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *StringStack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Pop remove and return top element of stack. Return false if stack is empty.
func (s *StringStack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
