package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Point struct{ x, y int }
type HeightMap map[Point]int

func readFile(name string) (HeightMap, int) {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var heightmap = make(HeightMap)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x := 0; x < len(line); x++ {
			val, _ := strconv.Atoi(string(line[x]))
			heightmap[Point{x, y}] = val
		}
		y++
	}

	return heightmap, y
}

func min(a, b int) int {
	if a < b { return a }
	return b
}

func lowestNeighbour(visited HeightMap, x, y int) int {
	lowest := math.MaxInt32
	left, ok := visited[Point{x - 1, y}]
	if ok { lowest = min(lowest, left) }
	right, ok := visited[Point{x + 1, y}]
	if ok { lowest = min(lowest, right) }
	up, ok := visited[Point{x, y - 1}]
	if ok { lowest = min(lowest, up) }
	down, ok := visited[Point{x, y + 1}]
	if ok { lowest = min(lowest, down) }
	return lowest
}

func riskFunction(heightmap HeightMap, x, y, size, mod int) int {
	rval := heightmap[Point{x % size, y % size}]
	for x >= size { rval++; x -= size }
	for y >= size { rval++; y -= size }
	for rval >= 10 { rval -= 9 }
	return rval
}

func one(heightmap HeightMap, size, mod int) int {
	visited := make(HeightMap)
	visited[Point{0, 0}] = 0
	unchanged := false
	iteration := 0
	for !unchanged {
		iteration++
		fmt.Println("Iteration: ", iteration)
		unchanged = true
		for y := 0; y <= size*mod; y++ {
			for x := 0; x <= size*mod; x++ {
				current, ok := visited[Point{x, y}]
				lowestN := lowestNeighbour(visited, x, y)
				prospective := lowestN + riskFunction(heightmap, x, y, size, mod)
				if !ok || (prospective < current) {
					visited[Point{x, y}] = prospective
					unchanged = false
				}
			}
		}
	}
	return visited[Point{size*mod -1, size*mod -1}]
}

func main() {
	heightmap, size := readFile("input")
	fmt.Println(one(heightmap, size, 1)) // 609
	fmt.Println(one(heightmap, size, 5)) // 2925
}
