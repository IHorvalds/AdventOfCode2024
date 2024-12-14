package main

import (
	"flag"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type stones []int

type key struct {
	stone  int
	blinks int
}

// mapping {stone, blinksLeft} to how many stones result after blinksLeft blinks starting from `stone`
var cache = map[key]int{}

func digitCount(i int) int {
	return int(math.Floor(math.Log10(float64(i)) + 1))
}

func splitDigits(i int) (int, int) {
	splitter := int(math.Pow10(digitCount(i) / 2))
	return i / splitter, i % splitter
}

func evenDigitCount(i int) bool {
	return digitCount(i)%2 == 0
}

// count of stones after blink b times starting from stone s
func stonesAfterBlinks(s int, bl int) int {
	if bl == 0 {
		return 1
	}

	k := key{s, bl}

	if res, ok := cache[k]; ok {
		return res
	}

	if s == 0 {
		cache[k] = stonesAfterBlinks(1, bl-1)
	} else if evenDigitCount(s) {
		a, b := splitDigits(s)
		cache[k] = stonesAfterBlinks(a, bl-1) + stonesAfterBlinks(b, bl-1)
	} else {
		cache[k] = stonesAfterBlinks(s*2024, bl-1)
	}

	return cache[k]
}

func parse(f string) (stones, error) {
	l, err := os.ReadFile(f)
	if err != nil {
		return stones{}, err
	}

	s := stones{}
	for _, b := range strings.Split(string(l), " ") {
		i, err := strconv.Atoi(b)
		if err != nil {
			return stones{}, err
		}
		s = append(s, i)
	}
	return s, nil
}

func main() {
	inputFlag := flag.String("input", "", "-input <input file>")
	blinks := flag.Int("blinks", 0, "-blinks <number of blinks>")

	flag.Parse()

	if inputFlag == nil {
		log.Fatal("empty input url")
	}

	if blinks == nil {
		log.Fatal("Need to blink")
	}

	st, err := parse(*inputFlag)
	if err != nil {
		log.Fatal(err)
	}

	c := 0
	for _, s := range st {
		c += stonesAfterBlinks(s, *blinks)
	}

	log.Printf("After %d blinks: %d", *blinks, c)
}
