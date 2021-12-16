package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var bits = map[rune]string{ //
	'0': "0000", '1': "0001", '2': "0010", '3': "0011", //
	'4': "0100", '5': "0101", '6': "0110", '7': "0111", //
	'8': "1000", '9': "1001", 'A': "1010", 'B': "1011", //
	'C': "1100", 'D': "1101", 'E': "1110", 'F': "1111"}

func readFile(name string) string {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

type Bits string

func (s *Bits) take(i int) string {
	t, u := (*s)[:i], (*s)[i:]
	*s = u
	return string(t)
}

func convert(input string) *Bits {
	rval := ""
	for _, c := range input { rval += bits[c] }
	bits := Bits(rval)
	return &bits
}

func binVal(input string) int {
	i, _ := strconv.ParseInt(input, 2, 64)
	return int(i)
}

func literal(input *Bits) int {
	bin := ""
	for {
		last, number := input.take(1), input.take(4)
		bin += number
		if last == "0" { break }
	}
	return binVal(bin)
}

type Packet struct {
	version, typeid, value int
	nested                 []Packet
}

func parse(input *Bits) Packet {
	version, typeid := binVal(input.take(3)), binVal(input.take(3))
	if typeid == 4 {
		return Packet{version, 4, literal(input), []Packet{}}
	}
	packet := Packet{version, typeid, 0, []Packet{}}
	lengthID := input.take(1)
	if lengthID == "0" {
		length := binVal(input.take(15))
		subs := Bits(input.take(length))
		for len(subs) > 0 { packet.nested = append(packet.nested, parse(&subs)) }
	} else {
		count := binVal(input.take(11))
		for i := 0; i < count; i++ { packet.nested = append(packet.nested, parse(input)) }
	}
	return packet
}

func allVersions(p Packet) int {
	rval := p.version
	for _, sp := range p.nested { rval += allVersions(sp) }
	return rval
}

func evaluate(p Packet) int {
	switch p.typeid {
	case 0: // sum
		rval := 0
		for _, p := range p.nested { rval += evaluate(p) }
		return rval
	case 1: // product
		rval := 1
		for _, p := range p.nested { rval *= evaluate(p) }
		return rval
	case 2: // minimum
		rval := math.MaxInt32
		for _, p := range p.nested {
			pval := evaluate(p)
			if pval < rval { rval = pval }
		}
		return rval
	case 3: // maximum
		rval := 0
		for _, p := range p.nested {
			pval := evaluate(p)
			if pval > rval { rval = pval }
		}
		return rval
	case 4: // value
		return p.value
	case 5: // greater
		a, b := evaluate(p.nested[0]), evaluate(p.nested[1])
		if a > b { return 1 }; return 0
	case 6: // less
		a, b := evaluate(p.nested[0]), evaluate(p.nested[1])
		if a < b { return 1 }; return 0
	case 7: // equal
		a, b := evaluate(p.nested[0]), evaluate(p.nested[1])
		if a == b { return 1 }; return 0
	default:
		panic("bad packet")
	}
}

func main() {
	puzzle := readFile("input")
	packet := parse(convert(puzzle))
	fmt.Println(allVersions(packet)) // 947
	fmt.Println(evaluate(packet))    // 660797830937
}
