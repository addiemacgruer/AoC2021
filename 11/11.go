package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct{ x, y int }
type Cavern map[Point]int
type StackType Point
type Stack []StackType

func (stack *Stack) push(b StackType) { *stack = append(*stack, b) }
func (stack *Stack) peek() StackType  { return (*stack)[len(*stack)-1] }
func (stack *Stack) isEmpty() bool    { return len(*stack) == 0 }
func (stack *Stack) pop() StackType {
	last := len(*stack) - 1
	rval := (*stack)[last]
	*stack = (*stack)[:last]
	return rval
}

func readFile(name string) Cavern {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var items = make(Cavern)
	y := 0
	for scanner.Scan() {
		item := scanner.Text()
		for x := 0; x < len(item); x++ {
			items[Point{x, y}], _ = strconv.Atoi(string(item[x]))
		}
		y++
	}

	return items
}

func (cavern Cavern) step() int {
	// increase the energy of each octopus, adding the ones which will flash to a stack
	var flashers = make(Stack, 0)
	for p, v := range cavern {
		next := v + 1
		if next > 9 {
			flashers.push(StackType(p))
		}
		cavern[p] = next
	}

	// for each one which will flash, increase the energy of the surrounding octopuses, keeping a set of the ones which have flashed already
	var flashed = make(map[Point]bool)
	for !flashers.isEmpty() {

		// have you already flashed?
		flash := Point(flashers.pop())
		if flashed[flash] {
			continue
		}

		// otherwise, increase the energy of surrounding octs, adding to the flashstack if needed
		flashed[flash] = true
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				p := Point{flash.x + x, flash.y + y}
				// OOB check
				if p.x < 0 || p.x >= 10 || p.y < 0 || p.y >= 10 {
					continue
				}
				next := cavern[p] + 1
				if next > 9 {
					flashers.push(StackType(p))
				}
				cavern[p] = next
			}
		}
	}

	for flash := range flashed {
		cavern[flash] = 0
	}
	return len(flashed)
}

func (cavern Cavern) getFlashes() (int, int) {
	flashes := 0
	step := 0
	hundred := 0
	for {
		step++
		flashcount := cavern.step()
		flashes += flashcount
		if step == 100 {
			hundred = flashes
		}
		if flashcount == 100 {
			return hundred, step
		}
	}
}

func main() {
	fmt.Println(readFile("input.test").getFlashes()) // 1656 195
	fmt.Println(readFile("input").getFlashes())      // 1640 312
}
