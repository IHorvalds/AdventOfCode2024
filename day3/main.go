package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Looking for "mul("
func findPrefix(s *string, start int) (bool, int) {
	if strings.HasPrefix((*s)[start:], "mul(") {
		return true, start + 3
	}

	return false, start
}

func findNumbers(s *string, start int) (int, int) {
	sub := (*s)[start:]
	r := regexp.MustCompile(`(\d|[1-9]\d{1,2}),(\d|[1-9]\d{1,2})\)`)
	ms := r.FindStringSubmatch(sub)
	if len(ms) != 3 {
		return 0, 0
	}
	if !strings.HasPrefix(sub, ms[0]) {
		return 0, 0
	}
	n1, err := strconv.Atoi(ms[1])
	if err != nil {
		return 0, 0
	}

	n2, err := strconv.Atoi(ms[2])
	if err != nil {
		return 0, 0
	}

	return n1, n2
}

func parseBytes(s string, enable bool) int {
	x := 0
	enabled := 1
	for i := 0; i < len(s); i++ {
		if enable {
			if strings.HasPrefix(s[i:], "do()") {
				enabled = 1
				i = i + 3
			} else if strings.HasPrefix(s[i:], "don't()") {
				enabled = 0
				i = i + 6
			}
		}

		if s[i] != 'm' {
			continue
		}

		found, i := findPrefix(&s, i)
		if !found {
			continue
		}

		n1, n2 := findNumbers(&s, i+1)
		x = x + enabled*n1*n2
	}
	return x
}

func main() {
	inputFlag := flag.String("input", "", "-input <input file>")

	flag.Parse()

	if inputFlag == nil {
		log.Fatal("empty input url")
	}

	bs, err := os.ReadFile(*inputFlag)
	if err != nil {
		log.Fatal(err)
	}

	res := parseBytes(string(bs), false)
	fmt.Printf("Part 1 result is: %d\n", res)
	res = parseBytes(string(bs), true)
	fmt.Printf("Part 2 result is: %d\n", res)
}
