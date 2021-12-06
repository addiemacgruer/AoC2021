package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FishMap map[int]int64

func readFile(name string) []int {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var fish []int
	scanner.Scan()
	fishString := scanner.Text()
	for _, i := range strings.Split(fishString, ",") {
		fint, _ := strconv.Atoi(i)
		fish = append(fish, fint)
	}
	return fish
}

func make_fishmap(raw []int) FishMap {
	var rval = make(FishMap)
	for _, i := range raw {
		rval[i]++
	}
	return rval
}

func countfish(fish FishMap) int64 {
	rval := int64(0)
	for _, i := range fish {
		rval += i
	}
	return rval
}

func next_day(fish FishMap) FishMap {
	var rval = make(FishMap)
	for k, v := range fish {
		if k == 0 {
			rval[6] += v
			rval[8] += v
		} else {
			rval[k-1] += v
		}
	}
	return rval
}

func main() {
	fishraw := readFile("input")
	fishmap := make_fishmap(fishraw)
	for day := 0; day < 256; day++ {
		fishmap = next_day(fishmap)
		if day == 80 {
			fmt.Println(countfish(fishmap)) // 380243
		}
	}
	fmt.Println(countfish(fishmap)) // 1708791884591
}
