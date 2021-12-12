package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const START = 0
const END = 1000

type Path struct{ start, end string }
type CaveList []int
type CaveMap map[int]CaveList
type PathList []int
type StackType PathList
type Stack []StackType
type Rejecter func(int, []int) bool

func (stack *Stack) push(b StackType) { *stack = append(*stack, b) }
func (stack *Stack) peek() StackType  { return (*stack)[len(*stack)-1] }
func (stack *Stack) isEmpty() bool    { return len(*stack) == 0 }
func (stack *Stack) pop() StackType {
	last := len(*stack) - 1
	rval := (*stack)[last]
	*stack = (*stack)[:last]
	return rval
}

func createPath(line string) Path {
	split := strings.Split(line, "-")
	return Path{split[0], split[1]}
}

func isSmall(cave string) bool {
	return cave[0] >= 'a' && cave[0] <= 'z'
}

func weight(cave string) int {
	if isSmall(cave) {
		return -1
	}
	return 1
}

func readFile(name string) CaveMap {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var caveMap = make(CaveMap)
	var names = make(map[string]int)
	names["start"] = START
	names["end"] = END
	next := 1
	for scanner.Scan() {
		path := createPath(scanner.Text())
		if _, exists := names[path.start]; !exists {
			names[path.start] = next * weight(path.start)
			next++
		}
		if _, exists := names[path.end]; !exists {
			names[path.end] = next * weight(path.end)
			next++
		}
		caveMap[names[path.start]] = append(caveMap[names[path.start]], names[path.end])
		caveMap[names[path.end]] = append(caveMap[names[path.end]], names[path.start])
	}
	return caveMap
}

func partOne(next int, work []int) bool {
	if next > 0 {
		return false
	}
	for _, w := range work {
		if w == next {
			return true
		}
	}
	return false
}

func partTwo(next int, work []int) bool {
	if next == START {
		return true
	}
	if !partOne(next, work) {
		return false
	}
	for i := 1; i < len(work); i++ {
		if work[i] > 0 {
			continue
		}
		for j := i + 1; j < len(work); j++ {
			if work[i] == work[j] {
				return true
			}
		}
	}
	return false
}

func pathCount(caveMap CaveMap, rejecter Rejecter) int {
	var workingPaths = make(Stack, 0)
	rval := 0
	workingPaths.push([]int{START})
	for !workingPaths.isEmpty() {
		work := workingPaths.pop()
		for _, next := range caveMap[work[len(work)-1]] {
			if rejecter(next, work) {
				continue
			}
			if next == END {
				rval++
			} else {
				prospective := make([]int, len(work))
				copy(prospective, work)
				prospective = append(prospective, next)
				workingPaths.push(prospective)
			}
		}
	}
	return rval
}

func main() {
	caveMap := readFile("input.test")
	fmt.Println(pathCount(caveMap, partOne)) // 10
	fmt.Println(pathCount(caveMap, partTwo)) // 36
	caveMap = readFile("input.test2")
	fmt.Println(pathCount(caveMap, partOne)) // 19
	fmt.Println(pathCount(caveMap, partTwo)) // 103
	caveMap = readFile("input.test3")
	fmt.Println(pathCount(caveMap, partOne)) // 226
	fmt.Println(pathCount(caveMap, partTwo)) // 3509
	caveMap = readFile("input")
	fmt.Println(pathCount(caveMap, partOne)) // 4304
	fmt.Println(pathCount(caveMap, partTwo)) // 118242
}
