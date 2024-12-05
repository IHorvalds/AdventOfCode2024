package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	xmas = "XMAS"
	samx = "SAMX"
)

func parse(f string) ([]string, error) {
	fd, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	v := []string{}
	r := bufio.NewReader(fd)
	for {
		l, err := r.ReadBytes('\n')
		if len(l) > 0 && l[len(l)-1] == '\n' {
			l = l[:len(l)-1]
		}
		if err == nil || errors.Is(err, io.EOF) {
			if len(l) > 0 {
				v = append(v, string(l))
			}
		}
		if err != nil {
			break
		}
	}
	return v, nil
}

func check(v *[]string, i, j int, k *[][]uint8) bool {
	return (*v)[i][j] == (*k)[0][0] && (*v)[i][j+2] == (*k)[0][2] && (*v)[i+1][j+1] == (*k)[1][1] && (*v)[i+2][j] == (*k)[2][0] && (*v)[i+2][j+2] == (*k)[2][2]
}

func part2(v *[]string) int {
	c := 0

	masSquares := [][][]uint8{
		{
			{'M', 0, 'S'},
			{0, 'A', 0},
			{'M', 0, 'S'},
		},
		{
			{'M', 0, 'M'},
			{0, 'A', 0},
			{'S', 0, 'S'},
		},
		{
			{'S', 0, 'M'},
			{0, 'A', 0},
			{'S', 0, 'M'},
		},
		{
			{'S', 0, 'S'},
			{0, 'A', 0},
			{'M', 0, 'M'},
		},
	}

	for i := range len(*v) - 2 {
		for j := range len((*v)[0]) - 2 {
			for k := range masSquares {
				if check(v, i, j, &masSquares[k]) {
					c++
					break
				}
			}
		}
	}

	return c
}

func countSquare(v *[]string, i, j int) int {
	c := 0

	copyVert := func(v *[]string, i, j int) string {
		b := []byte{}
		for x := range 4 {
			if i+x < len(*v) {
				b = append(b, (*v)[i+x][j])
			} else {
				b = append(b, '.')
			}
		}
		return string(b)
	}
	copyDiag := func(v *[]string, i, j, sign int) string {
		b := []byte{}
		for x := range 4 {
			if i+x < len(*v) && j+sign*x < len((*v)[0]) && j+sign*x >= 0 {
				b = append(b, (*v)[i+x][j+sign*x])
			} else {
				b = append(b, '.')
			}
		}
		return string(b)
	}

	if s := (*v)[i][j:]; strings.HasPrefix(s, xmas) || strings.HasPrefix(s, samx) {
		c++
	}

	if s := copyVert(v, i, j); s == xmas || s == samx {
		c++
	}
	if s := copyDiag(v, i, j, +1); s == xmas || s == samx {
		c++
	}
	if s := copyDiag(v, i, j+3, -1); s == xmas || s == samx {
		c++
	}

	return c
}

func part1(v *[]string) int {
	c := 0

	for i := range len(*v) {
		for j := range len((*v)[0]) {
			c = c + countSquare(v, i, j)
		}
	}

	return c
}

func main() {
	inputFlag := flag.String("input", "", "-input <input file>")

	flag.Parse()

	if inputFlag == nil {
		log.Fatal("empty input url")
	}
	v, err := parse(*inputFlag)
	if err != nil {
		log.Fatal(err)
	}
	count := part1(&v)
	fmt.Printf("Count is %d\n", count)
	count = part2(&v)
	fmt.Printf("Count 2 is %d\n", count)
}
