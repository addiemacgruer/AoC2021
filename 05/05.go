package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct{ x, y int }
type PointPair struct{ a, b Point }
type Field map[Point]int

func pointForString(item string) Point {
	bits := strings.Split(item, ",")
	x, _ := strconv.Atoi(bits[0])
	y, _ := strconv.Atoi(bits[1])
	return Point{x, y}
}

func pointPairForLine(item string) PointPair {
	bits := strings.Split(item, " ")
	a := pointForString(bits[0])
	b := pointForString(bits[2])
	return PointPair{a, b}
}

func readFile(name string) []PointPair {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var items []PointPair

	for scanner.Scan() {
		items = append(items, pointPairForLine(scanner.Text()))
	}

	return items
}

func twoPlusInField(field Field) int {
	count := 0
	for _, i := range field {
		if i > 1 {
			count++
		}
	}
	return count
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func increaseField(field Field, x, y int) {
	field[Point{x, y}] = field[Point{x, y}] + 1
}

func traceDiag(field Field, a Point, b Point, dx, dy int) {
	x, y := a.x, a.y
	for x >= min(a.x, b.x) && x <= max(a.x, b.x) {
		increaseField(field, x, y)
		x += dx
		y += dy
	}
}

func traceVent(p PointPair, field Field, diags bool) {
	if p.a.x == p.b.x { // vertical
		for y := min(p.a.y, p.b.y); y <= max(p.a.y, p.b.y); y++ {
			increaseField(field, p.a.x, y)
		}
	} else if p.a.y == p.b.y { // horizontal
		for x := min(p.a.x, p.b.x); x <= max(p.a.x, p.b.x); x++ {
			increaseField(field, x, p.a.y)
		}
	} else if diags {
		if p.a.x < p.b.x && p.a.y < p.b.y { // BL -> TR
			traceDiag(field, p.a, p.b, 1, 1)
		}
		if p.a.x > p.b.x && p.a.y < p.b.y { // BR -> TL
			traceDiag(field, p.a, p.b, -1, 1)
		}
		if p.a.x < p.b.x && p.a.y > p.b.y { // TL -> BR
			traceDiag(field, p.a, p.b, 1, -1)
		}
		if p.a.x > p.b.x && p.a.y > p.b.y { // TR -> BL
			traceDiag(field, p.a, p.b, -1, -1)
		}
	}
}

func tracePairs(points []PointPair, diags bool) int {
	var field = make(Field)
	for _, point := range points {
		traceVent(point, field, diags)
	}
	return twoPlusInField(field)
}

func main() {
	// text := readFile("input.test")
	text := readFile("input")
	fmt.Println(tracePairs(text, false)) // 4993
	fmt.Println(tracePairs(text, true))  // 21101
}
