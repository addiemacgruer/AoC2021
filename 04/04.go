package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Draws []int
type Board []int
type Boards = map[int]Board
type Drawn map[int]bool

func textToInts(text string, split string) []int {
	var rval []int
	drawsstrings := strings.Split(text, split)
	for _, s := range drawsstrings {
		if s == "" {
			continue
		}
		drawint, _ := strconv.Atoi(s)
		rval = append(rval, drawint)
	}
	return rval
}

func readFile(name string) (Draws, Boards) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal("failed to open")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	draws := textToInts(scanner.Text(), ",")
	scanner.Scan()

	boards := make(Boards)
	count := 0
	for scanner.Scan() {
		var board Board
		for i := 0; i < 5; i++ {
			line := textToInts(scanner.Text(), " ")
			for _, j := range line {
				board = append(board, j)
			}
			scanner.Scan()
		}
		boards[count] = board
		count++
	}

	return draws, boards
}

func wins(board Board, drawn Drawn) bool {
	for i := 0; i < 5; i++ {
		all := true
		for row := 0; row < 5; row++ {
			all = all && drawn[board[i*5+row]]
		}
		if all {
			return true
		}
		all = true
		for col := 0; col < 5; col++ {
			all = all && drawn[board[col*5+i]]
		}
		if all {
			return true
		}
	}
	return false
}

func score(board Board, drawn Drawn) int {
	rval := 0
	for _, v := range board {
		if !drawn[v] {
			rval += v
		}
	}
	return rval
}

func one(draws Draws, boards Boards) int {
	drawn := make(Drawn)
	for i, v := range draws {
		drawn[v] = true
		if i < 5 {
			continue
		}
		for _, board := range boards {
			if wins(board, drawn) {
				return score(board, drawn) * v
			}
		}
	}
	panic("no one won!")
}

func two(draws Draws, boards Boards) int {
	drawn := make(Drawn)
	for i, v := range draws {
		drawn[v] = true
		if i < 5 {
			continue
		}
		for i, board := range boards {
			if wins(board, drawn) {
				if len(boards) > 1 {
					delete(boards, i)
					continue
				}
				return score(board, drawn) * v
			}
		}
	}
	panic("no one lost!")
}

func main() {
	draws, boards := readFile("input")
	fmt.Println(one(draws, boards)) // 41668
	fmt.Println(two(draws, boards)) // 10478
}
