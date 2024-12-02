package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type safety int

const (
	SAFE safety = iota
	UNSAFE
)

func signum(i int) int {
	if i >= 0 {
		return 1
	}
	return -1
}

func removeFromSlice(s *[]int, i int) []int {
	v := []int{}
	for idx, val := range *s {
		if idx != i {
			v = append(v, val)
		}
	}
	return v
}

func isSafe(report *[]int) (safety, int) {
	if len(*report) <= 1 {
		return SAFE, len(*report)
	}

	delta := 0
	sign := signum((*report)[0] - (*report)[len(*report)-1]) // interval must be monotonically inc/dec
	prev := 0
	for idx, v := range *report {
		if idx == 0 {
			prev = v
			continue
		}

		delta = prev - v
		if delta == 0 {
			return UNSAFE, idx // no change
		}

		if signum(delta) != sign {
			return UNSAFE, idx // delta changed signs
		}

		if delta > 3 || delta < -3 {
			return UNSAFE, idx // too large jump
		}

		prev = v
	}

	return SAFE, len(*report)
}

// _REALLY_ naive implementation, but I can't think of a better one right now
func isSafeish(report *[]int) safety {
	r, top := isSafe(report)
	if r == SAFE {
		return SAFE
	}

	idx := top - 2
	if top < len(*report)-1 {
		top++
	}
	if idx < 0 {
		idx = 0
	}

	for ; idx <= top; idx++ {
		v := removeFromSlice(report, idx)
		r, _ := isSafe(&v)
		if r == SAFE {
			return SAFE
		}
	}
	return UNSAFE
}

func process(f string, proc func(*[]int) safety) (int, error) {
	fd, err := os.Open(f)
	if err != nil {
		return 0, err
	}
	defer fd.Close()

	r := bufio.NewReader(fd)
	s := 0
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return s, nil
			}
			return 0, err
		}

		v := []int{}
		for _, i := range strings.FieldsFunc(string(l), func(c rune) bool {
			return c == ' '
		}) {
			val, _ := strconv.Atoi(i)
			v = append(v, val)
		}
		if proc(&v) == SAFE {
			s++
		}
	}
}

func main() {
	inputFlag := flag.String("input", "", "-input <input file>")

	flag.Parse()

	if inputFlag == nil {
		log.Fatal("empty input url")
	}

	v, err := process(*inputFlag, func(s *[]int) safety {
		r, _ := isSafe(s)
		return r
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Safe: %d\n", v)

	v, err = process(*inputFlag, isSafeish)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Safeish: %d\n", v)
}
