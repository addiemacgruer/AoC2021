package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Subs map[string]string
type Step struct { pair string; level int }
type Dist map[byte]uint64

func readFile(name string) (string, Subs) {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	molecule := scanner.Text()
	scanner.Scan()
	var subs = make(Subs)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		subs[line[0]] = line[1]
	}
	return molecule, subs
}

func combined(a Dist, b Dist) Dist {
	var rval = make(Dist)
	for k, v := range a { rval[k] =  v }
	for k, v := range b { rval[k] += v }
	return rval
}

var cache = make(map[Step]Dist)

func evalutePair(pair string, subs Subs, limit int) Dist {
	next := Step{pair, limit}
	if existing, ok := cache[next]; ok { return existing }
	if limit == 0 {
		rval := Dist{pair[0]: 1}
		cache[next] = rval
		return rval
	}
	s, _ := subs[pair]
	mapA := evalutePair(string(pair[0])+s, subs, limit-1)
	mapB := evalutePair(s+string(pair[1]), subs, limit-1)
	rval := combined(mapA, mapB)
	cache[next] = rval
	return rval
}

func minMax(dist Dist) (uint64, uint64) {
	min, max := uint64(math.MaxUint64), uint64(0)
	for _, v := range dist {
		if v < min { min = v }
		if v > max { max = v }
	}
	return min, max
}

func evaluateMolecule(mol string, subs Subs, limit int) uint64 {
	var result = make(Dist)
	for i := 0; i < len(mol)-1; i++ {
		result = combined(result, evalutePair(mol[i:i+2], subs, limit))
	}
	result[mol[len(mol)-1]]++
	min, max := minMax(result)
	return max - min
}

func main() {
	mol, subs := readFile("input")
	fmt.Println(evaluateMolecule(mol, subs, 10)) // 2768
	fmt.Println(evaluateMolecule(mol, subs, 40)) // 2914365137499
}
