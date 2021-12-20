package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	PART_ONE = 2
	PART_TWO = 50
)

type Point struct{ x, y int }

type Floor map[Point]bool

type Image struct {
	floor Floor
	width int
	algo  string
	step  int
}

func (i Image) size() int {
	return len(i.floor)
}

func (i Image) at(p Point) int {
	rval := 0
	for y := p.y - 1; y <= p.y+1; y++ {
		for x := p.x - 1; x <= p.x+1; x++ {
			rval = rval << 1
			// horrible edge case for the 'real' puzzle
			if i.algo[0] == '#' {
				if x < -i.step || x >= i.width || y < -i.step || y >= i.width {
					if i.step%2 == 1 {
						rval += 1
						continue
					}
				}
			}
			if i.floor[Point{x, y}] {
				rval += 1
			}
		}
	}
	return rval
}

func (i Image) next() Image {
	n := Image{
		floor: make(Floor),
		width: i.width + 1,
		algo:  i.algo,
		step:  i.step + 1,
	}
	for y := -n.step; y < n.width; y++ {
		for x := -n.step; x < n.width; x++ {
			a := i.at(Point{x, y})
			if i.algo[a] == '#' {
				n.floor[Point{x, y}] = true
			}
		}
	}
	return n
}

func readImage(name string) Image {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	image := Image{}
	image.floor = make(Floor)

	scanner.Scan()
	image.algo = scanner.Text()
	if len(image.algo) != 512 { panic("bad algo") }

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { continue }
		for x := 0; x < len(line); x++ {
			if line[x] == '#' { image.floor[Point{x, y}] = true }
		}
		y++
	}
	image.width = y

	return image
}

func main() {
	image := readImage("input")
	enhance := image
	for i := 0; i < PART_ONE; i++ {
		enhance = enhance.next()
	}
	fmt.Println(enhance.size()) // 5225
	for i := PART_ONE; i < PART_TWO; i++ {
		enhance = enhance.next()
	}
	fmt.Println(enhance.size()) // 18131
}
