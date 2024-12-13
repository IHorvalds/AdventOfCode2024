package main

import (
	"bufio"
	"container/list"
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
)

type position struct {
	x, y int
}

type topographicMap struct {
	m          [][]int
	trailheads []position
}

func nextPositions(t *topographicMap, p position, positions *list.List) {
	if p.x > 0 && t.m[p.y][p.x-1] == t.m[p.y][p.x]+1 {
		positions.PushBack(position{p.x - 1, p.y})
	}

	if p.x < len(t.m[0])-1 && t.m[p.y][p.x+1] == t.m[p.y][p.x]+1 {
		positions.PushBack(position{p.x + 1, p.y})
	}

	if p.y > 0 && t.m[p.y-1][p.x] == t.m[p.y][p.x]+1 {
		positions.PushBack(position{p.x, p.y - 1})
	}

	if p.y < len(t.m)-1 && t.m[p.y+1][p.x] == t.m[p.y][p.x]+1 {
		positions.PushBack(position{p.x, p.y + 1})
	}
}

func (t *topographicMap) countTrails(start position, countAll bool) int {
	c := 0

	toVisit := list.New().Init()
	toVisit.PushBack(start)
	visited := map[position]struct{}{}
	for toVisit.Len() > 0 {
		v := toVisit.Remove(toVisit.Front()).(position)
		if countAll {
			if _, ok := visited[v]; ok {
				continue
			}

			visited[v] = struct{}{}
		}

		if t.m[v.y][v.x] == 9 {
			c++
			continue
		}

		nextPositions(t, v, toVisit)
	}

	return c
}

func part1(t *topographicMap) int {
	c := 0
	for _, th := range t.trailheads {
		c += t.countTrails(th, true)
	}
	return c
}

func part2(t *topographicMap) int {
	c := 0
	for _, th := range t.trailheads {
		c += t.countTrails(th, false)
	}
	return c
}

func parse(f string) (topographicMap, error) {
	fd, err := os.Open(f)
	if err != nil {
		return topographicMap{}, err
	}
	defer fd.Close()

	r := bufio.NewReader(fd)

	m := topographicMap{}

	line := 0
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return topographicMap{}, err
		}
		is := make([]int, len(l))
		for idx, b := range l {
			i, err := strconv.Atoi(string(b))
			if err != nil {
				return topographicMap{}, err
			}
			is[idx] = i
			if i == 0 {
				m.trailheads = append(m.trailheads, position{idx, line})
			}
		}
		m.m = append(m.m, is)
		line++
	}

	return m, nil
}

func main() {
	inputFlag := flag.String("input", "", "-input <input file>")

	flag.Parse()

	if inputFlag == nil {
		log.Fatal("empty input url")
	}

	m, err := parse(*inputFlag)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Trailhead score total: %d", part1(&m))
	log.Printf("Trailhead rating total: %d", part2(&m))
}
