package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type StackType byte
type Stack []StackType

func (stack *Stack) push(b StackType) { *stack = append(*stack, b) }
func (stack *Stack) peek() StackType  { return (*stack)[len(*stack)-1] }
func (stack *Stack) pop() StackType {
	last := len(*stack) - 1
	rval := (*stack)[last]
	*stack = (*stack)[:last]
	return rval
}

func readFile(name string) []string {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var items []string
	for scanner.Scan() {
		item := scanner.Text()
		items = append(items, item)
	}

	return items
}

func testCorruption(chunk string) int {
	var stack = make(Stack, 0)
	for i := 0; i < len(chunk); i++ {
		switch chunk[i] {
		case '{', '(', '[', '<': stack.push(StackType(chunk[i]))
		case ')': if stack.pop() != '(' { return 3 }
		case ']': if stack.pop() != '[' { return 57 }
		case '}': if stack.pop() != '{' { return 1197 }
		case '>': if stack.pop() != '<' { return 25137 }
		default: panic("bad character")
		}
	}

	rval := 0
	for len(stack) > 0 {
		rval *= 5
		switch stack.pop() {
		case '(': rval += 1
		case '[': rval += 2
		case '{': rval += 3
		case '<': rval += 4
		default: panic("bad character")
		}
	}
	return -rval
}

func checkAllSyntax(chunks []string) (int, int) {
	partOne := 0
	var partTwo = make([]int, 0)
	for _, chunk := range chunks {
		val := testCorruption(chunk)
		if val > 0 {
			partOne += val
		} else {
			partTwo = append(partTwo, -val)
		}
	}
	sort.Ints(partTwo)
	return partOne, partTwo[len(partTwo)/2]
}

func main() {
	chunks := readFile("input")
	a, b := checkAllSyntax(chunks)
	fmt.Println(a)
	fmt.Println(b)
}
