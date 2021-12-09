package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Point struct{ x, y int }
type HeightMap map[Point]int

func readFile(name string) HeightMap {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	y := 0
	var items = make(HeightMap)
	for scanner.Scan() {
		line := scanner.Text()
		for x := 0; x < len(line); x++ {
			items[Point{x, y}], _ = strconv.Atoi(string(line[x]))
		}
		y++
	}

	return items
}

func height(cavern HeightMap, x, y int) int {
	h, exists := cavern[Point{x, y}]
	if !exists {
		return 99
	}
	return h
}

func isLow(cavern HeightMap, x, y, h int) bool {
	up := height(cavern, x, y-1)
	down := height(cavern, x, y+1)
	left := height(cavern, x-1, y)
	right := height(cavern, x+1, y)
	if h >= up || h >= down || h >= left || h >= right {
		return false
	}
	return true
}

func consider(p Point, cavern HeightMap, basin map[Point]bool, prospective *[]Point) {
	if basin[p] {
		return
	}
	height, exists := cavern[p]
	if height == 9 || !exists {
		return
	}
	*prospective = append(*prospective, p)
}

func basinSize(cavern HeightMap, x, y int) int {
	basin := make(map[Point]bool)
	prospective := make([]Point, 0)
	prospective = append(prospective, Point{x, y})
	for len(prospective) > 0 {
		p := prospective[len(prospective)-1]
		prospective = prospective[:len(prospective)-1]
		basin[p] = true
		consider(Point{p.x, p.y - 1}, cavern, basin, &prospective)
		consider(Point{p.x, p.y + 1}, cavern, basin, &prospective)
		consider(Point{p.x - 1, p.y}, cavern, basin, &prospective)
		consider(Point{p.x + 1, p.y}, cavern, basin, &prospective)
	}
	return len(basin)
}

func one(cavern HeightMap) int {
	risk := 0
	for p, h := range cavern {
		if isLow(cavern, p.x, p.y, h) {
			risk += h + 1
		}
	}
	return risk
}

func two(cavern HeightMap) int {
	var basins = make([]int, 0)
	for p, h := range cavern {
		if isLow(cavern, p.x, p.y, h) {
			basins = append(basins, basinSize(cavern, p.x, p.y))
		}
	}
	sort.Ints(basins)
	return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
}

func main() {
	cavern := readFile("input")
	fmt.Println(one(cavern))
	fmt.Println(two(cavern))
}
