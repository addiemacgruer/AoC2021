package main

import (
	"fmt"
)

const (
	ROLLS = 3
	WIN1 = 1000
	WIN2 = 21
)

type State struct {
	a, b, as, bs int
	aturn        bool
}

type Wins struct { a, b uint64 }

var test = []int{4, 8}
var puzzle = []int{5, 8}
var cache map[State]Wins
var dist = []uint64{0, 0, 0, 1, 3, 6, 7, 6, 3, 1}

func mod(a, b int) int { for a > b { a -= b }; return a }

func one(start []int) int {
	a, b, as, bs, die := start[0], start[1], 0, 0, 0
	for {
		for r := 0; r < ROLLS; r++ {
			a += (die % 100) + 1
			die++
		}
		a = mod(a, 10)
		as += a
		if as >= WIN1 { return bs * die }
		for r := 0; r < ROLLS; r++ {
			b += (die % 100) + 1
			die++
		}
		b = mod(b, 10)
		bs += b
		if bs >= WIN1 { return as * die }
	}
}

func memo(s State, w Wins) Wins {
	cache[s] = w
	return w
}

func roll(s State) Wins {
	if cached, ok := cache[s]; ok { return cached }

	if s.as >= WIN2 { return memo(s,Wins{1,0}) }
	if s.bs >= WIN2 { return memo(s,Wins{0,1}) }

	w := Wins{}
	if s.aturn {
		for r := 3; r <= 9; r++ {
			next := mod(s.a+r, 10)
			a1 := roll(State{next, s.b, s.as + next, s.bs, false})
			w.a += a1.a * dist[r]
			w.b += a1.b * dist[r]
		}
	} else {
		for r := 3; r <= 9; r++ {
			next := mod(s.b+r, 10)
			b1 := roll(State{s.a, next, s.as, s.bs + next, true})
			w.a += b1.a * dist[r]
			w.b += b1.b * dist[r]
		}
	}
	return memo(s,w)
}

func two(start []int) uint64 {
	w := roll(State{start[0],start[1],0,0,true})
	if w.a > w.b { return w.a }
	return w.b
}

func main() {
	fmt.Println(one(test))   // 739785
	fmt.Println(one(puzzle)) // 1067724
	cache = make(map[State]Wins)
	fmt.Println(two(test))   // 444356092776315
	cache = make(map[State]Wins)
	fmt.Println(two(puzzle)) // 630947104784464
}
