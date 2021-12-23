package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct{ x, y, z int }
type Step struct {
	on       bool
	min, max Point
}
type Reactor map[Point]bool
type Reboot []Step

func (s Step) size() int64 {
	return int64(s.max.x+1-s.min.x) * int64(s.max.y+1-s.min.y) * int64(s.max.z+1-s.min.z)
}

func smaller(a, b int) int {
	if a > b { return b }; return a
}
func larger(a, b int) int {
	if a > b { return a }; return b
}
func smallBig(a, b, c, d int) (int, int) {
	return larger(a, b), smaller(c, d)
}

func overlap(a, b Step) *Step {
	if a.min.x > b.max.x || b.min.x > a.max.x { return nil }
	if a.min.y > b.max.y || b.min.y > a.max.y { return nil }
	if a.min.z > b.max.z || b.min.z > a.max.z { return nil }
	xmin, xmax := smallBig(a.min.x, b.min.x, a.max.x, b.max.x)
	ymin, ymax := smallBig(a.min.y, b.min.y, a.max.y, b.max.y)
	zmin, zmax := smallBig(a.min.z, b.min.z, a.max.z, b.max.z)
	return &Step{b.on, Point{xmin, ymin, zmin}, Point{xmax, ymax, zmax}}
}

func parseRange(line string) (int, int) {
	bits := strings.Split(line, "=")
	parts := strings.Split(bits[1], "..")
	min, _ := strconv.Atoi(parts[0])
	max, _ := strconv.Atoi(parts[1])
	return min, max
}

func stepForLine(line string) Step {
	bits := strings.Split(line, " ")
	on := bits[0] == "on"
	args := strings.Split(bits[1], ",")
	xmin, xmax := parseRange(args[0])
	ymin, ymax := parseRange(args[1])
	zmin, zmax := parseRange(args[2])
	return Step{on, Point{xmin, ymin, zmin}, Point{xmax, ymax, zmax}}
}

func readFile(name string) Reboot {
	file, _ := os.Open(name)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	steps := []Step{}
	for scanner.Scan() { steps = append(steps, stepForLine(scanner.Text())) }
	return steps
}

func c(val, limit int) int {
	if val <= -limit { return -limit }
	if val >= limit { return limit }
	return val
}

func onInCube(steps Reboot, limit int, cube Step) int64 {
	if limit < 0 { return 0 }
	olap := overlap(steps[limit], cube)
	var rval int64
	if olap == nil {
		rval = onInCube(steps, limit-1, cube)
	} else if steps[limit].on {
		rval = onInCube(steps, limit-1, cube) + olap.size() - onInCube(steps, limit-1, *olap)
	} else {
		rval = onInCube(steps, limit-1, cube) - onInCube(steps, limit-1, *olap)
	}
	return rval
}

func part2(steps Reboot) int64 {
	rval := int64(0)
	for i, step := range steps {
		if step.on {
			rval += step.size() - onInCube(steps, i-1, step)
		} else {
			rval -= onInCube(steps, i-1, step)
		}
	}
	return rval
}

func main() {
	text := readFile("input")

	fmt.Println(
			onInCube(text,
			len(text)-1,
			Step{false, Point{-50, -50, -50}, Point{50, 50, 50}})) // 545118

	fmt.Println(part2(text)) // 1227298136842375
}
