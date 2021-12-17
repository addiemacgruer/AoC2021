package main

import (
	"fmt"
	"math"
)

type Target struct{ xmin, xmax, ymin, ymax int }

var example = Target{20, 30, -10, -5}
var puzzle = Target{88, 125, -157, -103}

func max(a, b int) int { if a < b { return b }; return a }
func inRange(a, b, c int) bool { return a <= c && c <= b }

const (
	hit = iota
	stalled
	over
	under
)

func lob(dx, dy int, target Target) (int, int) {
	x, y, maxY := 0, 0, math.MinInt32
	for {
		x += dx
		y += dy
		if dx > 0 { dx-- }
		dy--
		maxY = max(y, maxY)
		if inRange(target.xmin, target.xmax, x) && inRange(target.ymin, target.ymax, y) {
			return hit, maxY
		}
		if dx == 0 && !inRange(target.xmin, target.xmax, x) { return stalled, 0 }
		if x > target.xmax { return over, 0 }
		if y < target.ymin { return under, 0 }
	}
}

func highestInitial(target Target) (int, int) {
	best, count := 0, 0
NextDX:
	for dx := 1; dx <= target.xmax; dx++ {
		hasHit := false
		for dy := target.ymin; dy < -target.ymin; dy++ {
			works, maxHeight := lob(dx, dy, target)
			switch works {
			case hit:
				best = max(best, maxHeight)
				count++
				hasHit = true
			case stalled:
				continue NextDX
			case over:
				if hasHit { continue NextDX }
			case under:
				// can't optimise - there might be mutliple 'blocks'
			}
		}
	}
	return best, count
}

func main() {
	fmt.Println(highestInitial(example)) // 45 112
	fmt.Println(highestInitial(puzzle))  // 12246 3528
}
