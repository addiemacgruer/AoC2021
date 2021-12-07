package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFile(name string) []int {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var items []int

	scanner.Scan()
	line := scanner.Text()
	for _, i := range strings.Split(line, ",") {
		item, _ := strconv.Atoi(i)
		items = append(items, item)
	}

	return items
}

func minmax(list []int) (int, int) {
	min, max := math.MaxInt32, 0
	for _, i := range list {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	return min, max
}

type scale func(int) int

func identity(val int) int {
	return val
}

func ladder(val int) int {
	return (val * (val + 1)) / 2
}

func ladder2(val int) int {
  rval := 0
  for i :=0; i <= val; i++ {
    rval += i
  }
  return rval
}

func fueltest(crabs []int, pos int, sf scale) int {
	rval := 0
	for _, c := range crabs {
		diff := pos - c
		if diff > 0 {
			rval += sf(diff)
		} else {
			rval += sf(-diff)
		}
	}
	return rval
}

func findLocalMin(crabs []int, sf scale) int {
	begin, end := minmax(crabs)
	lowest := math.MaxInt32
	for i := begin; i <= end; i++ {
		fuel := fueltest(crabs, i, sf)
		if fuel < lowest {
			lowest = fuel
		} else {
      return lowest
    }
	}
	return lowest
}

func main() {
	text := readFile("input")
	fmt.Println(findLocalMin(text, identity)) // 356992
	fmt.Println(findLocalMin(text, ladder2))   // 101268110
}
