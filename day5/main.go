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

// do better
type updatesList [][]int

func mid(v *[]int) (int, error) {
	if len(*v) == 0 {
		return 0, errors.New("empty slice")
	}
	if len(*v)%2 == 0 {
		return 0, errors.New("even number of elements")
	}
	return (*v)[len(*v)/2], nil
}

func parse(f string) (prioritySlice, updatesList, error) {
	fd, err := os.Open(f)
	if err != nil {
		return prioritySlice{}, updatesList{}, err
	}
	defer fd.Close()

	r := bufio.NewReader(fd)
	p := prioritySlice{}
	u := updatesList{}
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			return prioritySlice{}, updatesList{}, err
		}
		if len(l) == 0 {
			break
		}
		ns := strings.Split(string(l), "|")
		if len(ns) != 2 {
			panic("Invalid input")
		}

		i, err := strconv.Atoi(ns[0])
		if err != nil {
			return prioritySlice{}, updatesList{}, err
		}

		j, err := strconv.Atoi(ns[1])
		if err != nil {
			return prioritySlice{}, updatesList{}, err
		}
		p.insert(j, i)
	}
	p.sortPriorities()

	c := 0
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return prioritySlice{}, updatesList{}, err
		}
		v := []int{}
		for _, s := range strings.Split(string(l), ",") {
			i, err := strconv.Atoi(s)
			if err != nil {
				return prioritySlice{}, updatesList{}, err
			}
			v = append(v, i)
		}
		u = append(u, v)
		c++
	}

	return p, u, nil
}

func main() {
	inputFlag := flag.String("input", "", "-input <input file>")

	flag.Parse()

	if inputFlag == nil {
		log.Fatal("empty input url")
	}
	p, u, err := parse(*inputFlag)
	if err != nil {
		log.Fatal(err)
	}

	c := 0
	cc := 0
	for i := range u {
		if p.checkList(&u[i]) { // part1
			m, err := mid(&u[i])
			if err != nil {
				log.Fatal(err)
			}

			c = c + m
		} else { // part2
			p.sortList(&u[i])
			m, err := mid(&u[i])
			if err != nil {
				log.Fatal(err)
			}

			cc = cc + m
		}
	}

	fmt.Printf("Part 1: %d\n", c)
	fmt.Printf("Part 2: %d\n", cc)
}
