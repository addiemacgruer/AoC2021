package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct{ x, y int }

type SeaFloor struct {
	cucumbers     map[Point]bool
	width, height int
}

func makeSeafloor() SeaFloor {
	rval := SeaFloor{}
	rval.cucumbers = make(map[Point]bool)
	return rval
}

func (s SeaFloor) print() {
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			c, ok := s.cucumbers[Point{x, y}]
			switch {
			case !ok: fmt.Print(".")
			case c: fmt.Print(">")
			default: fmt.Print("v")
			}
		}
		fmt.Println()
	}
}

func (s SeaFloor) step() (SeaFloor, int) {
	first := makeSeafloor()
	first.width, first.height = s.width, s.height
	moves := 0

	for c, we := range s.cucumbers {
		if we == false {
			first.cucumbers[c] = we
			continue
		}
		next := Point{(c.x + 1) % s.width, c.y}
		if _, present := s.cucumbers[next]; present {
			first.cucumbers[c] = we
		} else {
			first.cucumbers[next] = we
			moves++
		}
	}

	second := makeSeafloor()
	second.width, second.height = s.width, s.height

	for c, we := range first.cucumbers {
		if we == true {
			second.cucumbers[c] = we
			continue
		}
		next := Point{c.x, (c.y + 1) % s.height}
		if _, present := first.cucumbers[next]; present {
			second.cucumbers[c] = we
		} else {
			second.cucumbers[next] = we
			moves++
		}
	}

	if len(s.cucumbers) != len(second.cucumbers) { panic("lost one!") }
	return second, moves
}

func readFile(name string) SeaFloor {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rval := makeSeafloor()
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		rval.width = len(line)
		for x := 0; x < len(line); x++ {
			if line[x] != '.' {
				rval.cucumbers[Point{x, y}] = line[x] == '>'
			}
		}
		y++
	}
	rval.height = y
	return rval
}

func main() {
	floor := readFile("input")
	step := 0
	var count int
	for {
		step++
		floor, count = floor.step()
		if count == 0 {
			break
		}
	}
	fmt.Println("Took", step, "steps")
}
