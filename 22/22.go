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
	for scanner.Scan() {
		steps = append(steps, stepForLine(scanner.Text()))
	}

	return steps
}

func c(val,limit int) int {
if (val <= -limit) { return -limit }
if (val >= limit) { return limit }
return val
}

func execute(steps Reboot,limit int) uint64 {
	reactor := make(Reactor)
	for _, step := range steps {
		fmt.Print(".")
		if step.min.x > limit || step.min.y > limit || step.min.z > limit {continue}
		if step.max.x < -limit || step.max.y < -limit || step.max.z < -limit {continue}
		for x := c(step.min.x,limit); x <= c(step.max.x,limit); x++ {
			for y := c(step.min.y,limit); y <= c(step.max.y,limit); y++ {
				for z := c(step.min.z,limit); z <= c(step.max.z,limit); z++ {
					if step.on {
						reactor[Point{x, y, z}] = true
					} else {
						delete(reactor, Point{x, y, z})
					}
				}
			}
		}
	}
	return uint64(len(reactor))
}

func main() {
	text := readFile("test3")
	fmt.Println(text)
	fmt.Println(execute(text,50)) // 545118
}
