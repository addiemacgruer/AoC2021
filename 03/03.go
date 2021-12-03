package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(name string) []string {
	file, err := os.Open(name)
	if err != nil {
		panic("failed to open: " + name)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var diags []string
	for scanner.Scan() {
		line := scanner.Text()
		diags = append(diags, line)
	}
	return diags
}

func countOnes(bit int, diags []string) int {
	ones := 0
	for _, diag := range diags {
		if diag[bit] == '1' {
			ones++
		}
	}
	return ones
}

func ternary(cond bool, a int, b int) int {
  if cond {
    return a
  } else {
    return b
  }
}

func one(diags []string) int {
	gamma, epsilon := 0, 0
	for bit := 0; bit < len(diags[0]); bit++ {
		onec := countOnes(bit, diags) > (len(diags) / 2)
		gamma = 2 * gamma + ternary(onec,1,0)
		epsilon = 2 * epsilon + ternary(onec,0,1)
	}
	return gamma * epsilon
}

func filter(bit int, diags []string, a byte, b byte) []string {
	onec := countOnes(bit, diags)
	var target byte
  if onec < len(diags) - onec {
    target = a
  } else {
    target = b
  }

	var result []string
	for _, diag := range diags {
		if diag[bit] == target {
			result = append(result, diag)
		}
	}
	return result
}

func machinefilter(diags[]string, a byte, b byte) string {
	for bit := 0; bit < len(diags[0]); bit++ {
		diags = filter(bit, diags, a, b)
		if len(diags) == 1 {
			return diags[0]
		}
	}
	panic("Couldn't narrow down")
}

func two(diags []string) int64 {
  oxy,_ := strconv.ParseInt(machinefilter(diags, '0', '1'), 2, 64)
  co2,_ := strconv.ParseInt(machinefilter(diags, '1', '0'), 2, 64)
	return oxy * co2
}

func main() {
	diags := readFile("input.test")
	fmt.Println(one(diags)) // 198
	fmt.Println(two(diags)) // 230
	diags = readFile("input")
	fmt.Println(one(diags)) // 1997414
	fmt.Println(two(diags)) // 1032597
}
