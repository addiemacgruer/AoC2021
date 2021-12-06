package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(name string) []int {
	file, _ := os.Open(name)
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

func main() {
	text := readFile("input.test")
  fmt.Println(text)
}
