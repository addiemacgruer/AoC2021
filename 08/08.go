package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
	"strings"
)

type Entry struct {
	patterns []string
	output []string
}

type ByLength []string
func (a ByLength) Len() int           { return len(a) }
func (a ByLength) Less(i, j int) bool { return len(a[i]) < len(a[j]) }
func (a ByLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func entryForLine(line string) Entry {
	l := strings.Split(line, " | ")
	dig := strings.Split(l[0], " ")
	out := strings.Split(l[1], " ")
	return Entry{dig, out}
}

func readFile(name string) []Entry {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var entries []Entry
	for scanner.Scan() {
		entries = append(entries, entryForLine(scanner.Text()))
	}
	return entries
}

func one(entries []Entry) int {
	rval := 0
	for _, entry := range entries {
		for _, out := range entry.output {
			switch len(out) {
			case 2, 3, 4, 7:
				rval++
			}
		}
	}
	return rval
}

func byteForDigit(input string) byte {
	rval := byte(0)
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case 'a':
			rval = rval | byte(0b00000001)
		case 'b':
			rval = rval | byte(0b00000010)
		case 'c':
			rval = rval | byte(0b00000100)
		case 'd':
			rval = rval | byte(0b00001000)
		case 'e':
			rval = rval | byte(0b00010000)
		case 'f':
			rval = rval | byte(0b00100000)
		case 'g':
			rval = rval | byte(0b01000000)
		}
	}
	return rval
}

// given 2, 3, 5
func findThree(one, x, y, z byte) (byte, byte, byte) {
	if one&x == one {
		return x, y, z
	}
	if one&y == one {
		return y, x, z
	}
	if one&z == one {
		return z, x, y
	}
	panic("None of them!")
}

// given 0, 6, 9
func findZero(D, x, y, z byte) (byte, byte, byte) {
	if D&x == 0 {
		return x, y, z
	}
	if D&y == 0 {
		return y, x, z
	}
	if D&z == 0 {
		return z, x, y
	}
	panic("None of them!")
}

func decode(digs []string) []byte {
	sort.Sort(ByLength(digs))

	r := make([]byte, 10)
	var x, y, m, n byte

	r[1] = byteForDigit(digs[0])
	r[7] = byteForDigit(digs[1])
	r[4] = byteForDigit(digs[2])
	r[8] = byteForDigit(digs[9])

	r[3], x, y = findThree(r[1],
		byteForDigit(digs[3]),
		byteForDigit(digs[4]),
		byteForDigit(digs[5]))

	if bits.OnesCount8(r[4]&x) == 2 {
		r[2], r[5] = x, y
	} else {
		r[2], r[5] = y, x
	}

	D := r[2] & r[3] & r[4] & r[5]

	r[0], m, n = findZero(D,
		byteForDigit(digs[6]),
		byteForDigit(digs[7]),
		byteForDigit(digs[8]))

	if m&r[1] == r[1] {
		r[6], r[9] = n, m
	} else {
		r[6], r[9] = m, n
	}

	return r
}

func valForDig(dec []byte, inp byte) int {
	for i, d := range dec {
		if d == inp {
			return i
		}
	}
	panic("oops!")
}

func two(entries []Entry) int {
	rval := 0
	for _, entry := range entries {
		dec := decode(entry.patterns)
		ans := 0
		for _, d := range entry.output {
			ans *= 10
			ans += valForDig(dec, byteForDigit(d))
		}
		rval += ans
	}
	return rval
}

func main() {
	entries := readFile("input")
	fmt.Println(one(entries)) // 245
	fmt.Println(two(entries)) // 983026
}
