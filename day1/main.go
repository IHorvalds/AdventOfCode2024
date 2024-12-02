package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Problem text
// https://adventofcode.com/2024/day/1

func getInput(f string) (*[2]AVL, error) {
	b, e := os.ReadFile(f)
	if e != nil {
		return nil, e
	}

	avls := [2]AVL{}
	for _, line := range strings.Split(string(b), "\n") {
		for idx, i := range strings.FieldsFunc(line, func(c rune) bool {
			return c == ' '
		}) {
			val, _ := strconv.Atoi(i)
			avls[idx].insert(val)
		}
	}

	return &avls, nil
}

func part1(v *[2]AVL) int {
	return diff(&v[0], &v[1])
}

func part2(v *[2]AVL) int {
	freq := map[int]int{}
	for i := range v[1].traverse() {
		if _, ok := freq[i]; !ok {
			freq[i] = 0
		}
		freq[i] = freq[i] + 1
	}

	sim := 0
	for i := range v[0].traverse() {
		if f, ok := freq[i]; ok {
			sim = sim + i*f
		}
	}
	return sim
}

func main() {
	inputFlag := flag.String("input", "", "-input <input file>")

	flag.Parse()

	if inputFlag == nil {
		log.Fatal("empty input url")
	}

	v, err := getInput(*inputFlag)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Diff is: %d\n", part1(v))
	fmt.Printf("Similarity is: %d\n", part2(v))
}
