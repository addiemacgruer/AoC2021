package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct{ x, y int }
type Fold struct {
	axis  string
	value int
}
type PointMap map[Point]bool

func (pm PointMap) size() int {
	return len(pm)
}

func (pm PointMap) print() {
	for y := 0; y < 10; y++ {
		for x := 0; x < 50; x++ {
			if pm[Point{x, y}] {
				fmt.Print("x")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func readFile(name string) (PointMap, []Fold) {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var points = make(PointMap)
	var folds = make([]Fold, 0)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		if strings.HasPrefix(text, "fold") {
			val, _ := strconv.Atoi(text[13:])
			fold := Fold{string(text[11]), val}
			folds = append(folds, fold)
		} else {
			split := strings.Split(text, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			points[Point{x, y}] = true
		}
	}

	return points, folds
}

func doFolds(points PointMap, folds []Fold, limit int) int {
	for i, f := range folds {
		var nextPoints = make(PointMap)
		for p := range points {
			switch f.axis {
			case "x":
				if p.x == f.value {
					continue
				}
				nextx := p.x
				if nextx > f.value {
					nextx = 2*f.value - nextx
				}
				nextPoints[Point{nextx, p.y}] = true
			case "y":
				if p.y == f.value {
					continue
				}
				nexty := p.y
				if nexty > f.value {
					nexty = 2*f.value - nexty
				}
				nextPoints[Point{p.x, nexty}] = true
			default:
				panic("bad fold!")
			}
		}
		points = nextPoints
		if i == limit {
			break
		}
	}
	if limit == -1 {
		points.print()
	}
	return points.size()
}

func main() {
	points, folds := readFile("input")
	fmt.Println(doFolds(points, folds, 0)) // 592
	doFolds(points, folds, -1)             // JGAJEFKU
}
