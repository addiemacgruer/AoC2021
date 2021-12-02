package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	forward = iota
	down
	up
)

type Command struct {
	dir  int
	dist int
}

func commandForLine(dir int, line string) Command {
	dist, _ := strconv.Atoi(line[strings.Index(line, " ")+1:])
	return Command{dir, dist}
}

func readFile(name string) []Command {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal("failed to open")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var commands []Command

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "forward"):
			commands = append(commands, commandForLine(forward, line))
		case strings.HasPrefix(line, "down"):
			commands = append(commands, commandForLine(down, line))
		case strings.HasPrefix(line, "up"):
			commands = append(commands, commandForLine(up, line))
		default:
			log.Fatal("didn't understand ", line)
		}
	}

	return commands
}

func one(commands []Command) {
	horizontal, depth := 0, 0
	for _, command := range commands {
		switch command.dir {
		case forward:
			horizontal += command.dist
		case up:
			depth -= command.dist
		case down:
			depth += command.dist
		}
	}
	fmt.Println(horizontal, "x", depth, "=", horizontal*depth)
}

func two(commands []Command) {
	horizontal, depth, aim := 0, 0, 0
	for _, command := range commands {
		switch command.dir {
		case forward:
			horizontal += command.dist
			depth += (command.dist * aim)
		case up:
			aim -= command.dist
		case down:
			aim += command.dist
		}
	}
	fmt.Println(horizontal, "x", depth, "=", horizontal*depth)
}

func main() {
	commands := readFile("input")
	one(commands) // 1604850
	two(commands) // 1685186100
}
