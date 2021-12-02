package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(name string) []int {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal("failed to open")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var items []int

	for scanner.Scan() {
		item, _ := strconv.Atoi(scanner.Text())
		items = append(items, item)
	}

	return items
}

func interval(gaps []int, gap int) int {
	greater := 0
	for i := 0; i < len(gaps)-gap; i++ {
		if gaps[i+gap] > gaps[i] {
			greater++
		}
	}
	return greater
}

func main() {
	text := readFile("input")
	fmt.Println(interval(text, 1)) // 1548
	fmt.Println(interval(text, 3)) // 1589
}
