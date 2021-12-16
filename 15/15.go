package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Point struct{ x, y int }
type HeightMap [][]int

func createHeightMap(size int) HeightMap {
	var heightmap = make([][]int,size+1)
	for y := 0; y <= size; y++ {
		heightmap[y] = make([]int,size+1)
		for x := 0; x <= size; x++ {
			heightmap[y][x] = math.MaxInt32
		}
	}
	return heightmap
}

func (h HeightMap) at(x,y int) int {
	if (y < 0 || y >= len(h) || x < 0 || x >= len(h[y])) { return math.MaxInt32 }
	return h[y][x]
}

func (h HeightMap) set(x,y,z int) {
	h[y][x] = z
}

func readFile(name string) (HeightMap, int) {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var heightmap HeightMap
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if y == 0 {heightmap = createHeightMap(len(line))}
		for x := 0; x < len(line); x++ {
			val, _ := strconv.Atoi(string(line[x]))
			heightmap.set(x,y,val)
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
	z := visited.at(x - 1, y)
	z = min(z, visited.at(x + 1, y))
	z = min(z, visited.at(x, y - 1))
	return min(z, visited.at(x, y + 1))
}

func riskFunction(heightmap HeightMap, x, y, size, mod int) int {
	rval := heightmap.at(x % size, y % size)
	for x >= size { rval++; x -= size }
	for y >= size { rval++; y -= size }
	for rval >= 10 { rval -= 9 }
	return rval
}

func expanded(basic HeightMap, size, mod int) HeightMap {
	rval := createHeightMap(size*mod)
	for y := 0; y < size*mod; y++ {
	for x := 0; x < size*mod; x++ {
		rval.set(x,y,riskFunction(basic,x,y,size,mod))
	}
	}
	return rval
}

func one(basic HeightMap, size, mod int) int {
	visited := createHeightMap(size*mod)
	heightmap := expanded(basic, size, mod)
	visited.set(0,0,0)
	unchanged := false
	for !unchanged {
		unchanged = true
		for y := 0; y <= size*mod; y++ {
			for x := 0; x <= size*mod; x++ {
				current := visited.at(x, y)
				lowestN := lowestNeighbour(visited, x, y)
				prospective := lowestN + heightmap.at(x, y)
				if prospective < current {
					visited.set(x,y,prospective)
					unchanged = false
				}
			}
		}
	}
	return visited.at(size*mod -1, size*mod -1)
}

func main() {
	theightmap, tsize := readFile("input.test")
	fmt.Println(one(theightmap, tsize, 1)) // 40
	fmt.Println(one(theightmap, tsize, 5)) // 315
	heightmap, size := readFile("input")
	fmt.Println(one(heightmap, size, 1)) // 609
	fmt.Println(one(heightmap, size, 5)) // 2925
}
