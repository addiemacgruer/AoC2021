package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// pair either has a leftvalue or a pointer to one (lval or lp), likewise right
// every pair but the root has a pointer to its superpair, s
type Pair struct {
	lval, rval int
	lp, rp, s  *Pair
}

type String string
func (s *String) take(i int) string {
	t, u := (*s)[:i], (*s)[i:]
	*s = u
	return string(t)
}

func parseLine(line *String) *Pair {
	pair := Pair{}
	pair.s = nil

	if line.take(1) != "[" { panic("bad [!") }

	if (*line)[0] == '[' {
		pair.lp = parseLine(line)
		pair.lp.s = &pair
	} else {
		v, _ := strconv.Atoi(line.take(1))
		pair.lval = v
	}

	if line.take(1) != "," { panic("bad ,!") }

	if (*line)[0] == '[' {
		pair.rp = parseLine(line)
		pair.rp.s = &pair
	} else {
		v, _ := strconv.Atoi(line.take(1))
		pair.rval = v
	}

	if line.take(1) != "]" { panic("bad ]!") }

	return &pair
}

func parse(s string) *Pair {
	ss := String(s)
	return parseLine(&ss)
}

func readFile(name string) []string {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var items []string
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}

	return items
}

func (p *Pair) addleft(i int) {
	if p.s == nil {
		return
	}
	if p.s.lp != nil && p.s.lp != p {
		rt := p.s.lp
		for rt.rp != nil {
			rt = rt.rp
		}
		rt.rval += i
	} else if p.s.lp != nil {
		p.s.addleft(i)
	} else {
		p.s.lval += i
	}
}

func (p *Pair) addright(i int) {
	if p.s == nil {
		return
	}
	if p.s.rp != nil && p.s.rp != p {
		lt := p.s.rp
		for lt.lp != nil {
			lt = lt.lp
		}
		lt.lval += i
	} else if p.s.rp != nil {
		p.s.addright(i)
	} else {
		p.s.rval += i
	}
}

func explode(p *Pair, side bool, level int) (*Pair, bool) {
	if p.lp != nil {
		_, b := explode(p.lp, false, level+1)
		if b { return p,true }
	}
	if p.rp != nil {
		_, b := explode(p.rp, true, level+1)
		if b { return p,true }
	}
	if level > 4 {
		p.addleft(p.lval)
		p.addright(p.rval)
		if !side {
			p.s.lval = 0
			p.s.lp = nil
		} else {
			p.s.rval = 0
			p.s.rp = nil
		}
		return nil, true
	}
	return p, false
}

func split(p *Pair) (*Pair, bool) {
	if p.lp != nil {
		_, d := split(p.lp)
		if d { return p,true }
	} else {
		if p.lval >= 10 {
			split := Pair{}
			split.lval = p.lval / 2
			split.rval = p.lval - split.lval
			split.s = p
			p.lval = 0
			p.lp = &split
			return p, true
		}
	}
	if p.rp != nil {
		_, d := split(p.rp)
		if d { return p,true }
	} else {
		if p.rval >= 10 {
			split := Pair{}
			split.lval = p.rval / 2
			split.rval = p.rval - split.lval
			split.s = p
			p.rval = 0
			p.rp = &split
			return p, true
		}
	}
	return p,false
}

func add(a, b *Pair) *Pair {
	p := Pair{}
	p.lp = a
	p.lp.s = &p
	p.rp = b
	p.rp.s = &p
	q := &p
Exploding:
	for {
		did := true
		for did {
			q, did = explode(q, false, 1)
			if did { continue Exploding }
		}
		did = true
		for did {
			q, did = split(q)
			if did { continue Exploding }
		}
		break
	}
	return q
}

func (p Pair) magnitude() int {
	rval := 0
	if p.lp != nil { rval += 3 * p.lp.magnitude() } else { rval += 3 * p.lval }
	if p.rp != nil { rval += 2 * p.rp.magnitude() } else { rval += 2 * p.rval }
	return rval
}

func part1(text []string) {
	ans := parse(text[0])
	for i := 1; i < len(text); i++ {
		plus := parse(text[i])
		ans = add(ans, plus)
	}
	fmt.Println(ans.magnitude())
}

func part2(text []string) {
	max := 0
	for a := 0; a < len(text); a++ {
		for b := 0; b < len(text); b++ {
			if a == b {
				continue
			}
			first := parse(text[a])
			second := parse(text[b])
			sum := add(first, second)
			mag := sum.magnitude()
			if mag > max {
				max = mag
			}
		}
	}
	fmt.Println(max)
}

func main() {
	text := readFile("input.test")
	part1(text)
	part2(text)
	text = readFile("input")
	part1(text)
	part2(text)
}
