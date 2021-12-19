package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

const FACES = 24
const ENOUGH = 12

type Point struct{ x, y, z int }
type void struct{}

var member void

func (a Point) hash() int {
	return a.x<<16 | a.y<<8 | a.z
}

func (a Point) plus(b Point) Point {
	return Point{a.x + b.x, a.y + b.y, a.z + b.z}
}

func (a Point) minus(b Point) Point {
	return Point{a.x - b.x, a.y - b.y, a.z - b.z}
}

var mu sync.Mutex

type Scanner struct {
	pos, cpos   Point
	face, cface int
	locked      bool
	beacons     []Point
	adj         []Point
	knownBad    map[int]bool
}

func (s Scanner) size() int { return len(s.beacons) }

func (s Scanner) adjust(i int) Point {
	twist, spin := s.face%6, s.face/6
	raw := s.beacons[i]
	var twisted Point
	switch twist {
	case 0: twisted = raw
	case 1: twisted = Point{raw.y, -raw.x, raw.z}
	case 2: twisted = Point{-raw.x, -raw.y, raw.z}
	case 3: twisted = Point{-raw.y, raw.x, raw.z}
	case 4: twisted = Point{-raw.z, raw.y, raw.x}
	case 5: twisted = Point{raw.z, raw.y, -raw.x}
	default: panic("bad twist")
	}
	var spun Point
	switch spin {
	case 0: spun = twisted
	case 1: spun = Point{twisted.x, -twisted.z, twisted.y}
	case 2: spun = Point{twisted.x, -twisted.y, -twisted.z}
	case 3: spun = Point{twisted.x, twisted.z, -twisted.y}
	default: panic("bad spin")
	}
	return spun.plus(s.pos)
}

func (s *Scanner) at(i int) Point {
	if s.pos != s.cpos || s.face != s.cface {
		s.adj = make([]Point, len(s.beacons))
		for i := 0; i < len(s.beacons); i++ {
			s.adj[i] = s.adjust(i)
		}
		s.cpos = s.pos
		s.cface = s.face
	}
	return s.adj[i]
}

func atoi(s string) int {
	rval, err := strconv.Atoi(s)
	if err != nil {
		panic("bad int!")
	}
	return rval
}

func readFile(name string) []Scanner {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var items []Scanner
	var wip *Scanner
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "---") {
			if wip != nil {
				items = append(items, *wip)
			}
			wip = &Scanner{}
			wip.knownBad = make(map[int]bool)
			wip.cface = -1
			continue
		}

		split := strings.Split(line, ",")
		p := Point{atoi(split[0]), atoi(split[1]), atoi(split[2])}
		(*wip).beacons = append((*wip).beacons, p)
	}
	items = append(items, *wip)

	items[0].locked = true
	return items
}

func testOverlap(a Scanner, b Scanner) int {
	both := a.size() + b.size()
	var over = make(map[int]void, both)
	ae := a.size()
	for i := 0; i < ae; i++ {
		over[a.at(i).hash()] = member
	}
	be := b.size()
	for i := 0; i < be; i++ {
		over[b.at(i).hash()] = member
	}
	return both - len(over)
}

func findfit(lock *Scanner, unlock *Scanner) bool {
	for face := 0; face < FACES; face++ {
		(*unlock).face = face
		for i := 0; i < (*lock).size(); i++ {
			for j := 0; j < (*unlock).size(); j++ {
				(*unlock).pos = Point{0, 0, 0}
				p := (*unlock).at(j)
				(*unlock).pos = (*lock).at(i).minus(p)
				count := testOverlap(*lock, *unlock)
				if count >= ENOUGH {
					return true
				}
			}
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattan(a, b Point) int {
	return abs(a.x-b.x) + abs(a.y-b.y) + abs(a.z-b.z)
}

func checkPair(text []Scanner, i, j int, out chan bool) {
	if !text[i].locked || text[j].locked || i == j || text[i].knownBad[j] {
		out <- false
		return
	}
	if findfit(&(text[i]), &(text[j])) {
		text[j].locked = true
		fmt.Println("Locked", i, "vs", j, "@", text[j].face, text[j].pos)
		out <- true
	} else {
		mu.Lock()
		text[i].knownBad[j] = true
		mu.Unlock()
	}
	out <- false
}

func main() {
	text := readFile("input")
	for {
		testing := false
		for i := 0; i < len(text); i++ {
			var channels = make([]chan bool, len(text))
			for j := 0; j < len(text); j++ {
				channels[j] = make(chan bool)
				go checkPair(text, i, j, channels[j])
			}
			for j := 0; j < len(text); j++ {
				if <- channels[j] {
					testing = true
				}
			}
		}
		if !testing {
			break
		}
	}
	unique := make(map[Point]void)
	for _, t := range text {
		for i := 0; i < t.size(); i++ {
			unique[t.at(i)] = member
		}
	}
	fmt.Println(len(unique))

	maxdist := 0
	for i := 0; i < len(text); i++ {
		for j := i + 1; j < len(text); j++ {
			mh := manhattan(text[i].pos, text[j].pos)
			if mh > maxdist {
				maxdist = mh
			}
		}
	}
	fmt.Println(maxdist)
}
