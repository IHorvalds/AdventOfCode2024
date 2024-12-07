package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"slices"
)

const (
	obstacle = '#'
	visited  = 'x'
	newPos   = '.'
)

type floorPlan [][]byte

func (f floorPlan) outOfBounds(p position) bool {
	return p.x < 0 || p.x >= len(f[0]) || p.y < 0 || p.y >= len(f)
}

// the state of the guard when they see an obstacle
type states []guard

type direction int

const ( // order is important for direction updates
	UP direction = iota
	RIGHT
	DOWN
	LEFT
)

type position struct {
	x, y int
}

type guard struct {
	d direction
	p position
}

func (g *guard) walk(m *floorPlan) int {
	nextPos := g.p
	unique := 1
	for {
		switch g.d {
		case UP:
			nextPos.y--
		case DOWN:
			nextPos.y++
		case LEFT:
			nextPos.x--
		case RIGHT:
			nextPos.x++
		}
		if m.outOfBounds(nextPos) {
			break
		}
		if (*m)[nextPos.y][nextPos.x] == newPos {
			unique++
			(*m)[nextPos.y][nextPos.x] = visited
		}
		if (*m)[nextPos.y][nextPos.x] == obstacle {
			g.updateDirection()
			nextPos = g.p
		} else {
			g.p = nextPos
		}
	}
	return unique
}

func (g *guard) walk2(m *floorPlan) int {
	nextPos := g.p
	c := 0
	for {
		switch g.d {
		case UP:
			nextPos.y--
		case DOWN:
			nextPos.y++
		case LEFT:
			nextPos.x--
		case RIGHT:
			nextPos.x++
		}
		if m.outOfBounds(nextPos) {
			break
		}
		if (*m)[nextPos.y][nextPos.x] == newPos {
			g2 := *g
			if g2.checkLoop(m, nextPos) {
				c++
			}
			(*m)[nextPos.y][nextPos.x] = visited
		}
		if (*m)[nextPos.y][nextPos.x] == obstacle {
			g.updateDirection()
			nextPos = g.p
		} else {
			g.p = nextPos
		}
	}
	return c
}

func (g *guard) checkLoop(m *floorPlan, obs position) bool {
	nextPos := g.p
	st := states{*g}
	g.updateDirection()
	ov := (*m)[obs.y][obs.x]
	(*m)[obs.y][obs.x] = obstacle
	defer func() {
		(*m)[obs.y][obs.x] = ov
	}()
	for {
		switch g.d {
		case UP:
			nextPos.y--
		case DOWN:
			nextPos.y++
		case LEFT:
			nextPos.x--
		case RIGHT:
			nextPos.x++
		}
		if m.outOfBounds(nextPos) {
			return false
		}
		if (*m)[nextPos.y][nextPos.x] == obstacle {
			if slices.Contains(st, *g) {
				return true
			}
			st = append(st, *g)
			g.updateDirection()
			nextPos = g.p
		} else {
			g.p = nextPos
		}
	}
}

func (g *guard) updateDirection() {
	g.d = (g.d + 1) % 4
}

func parse(f string) (*floorPlan, *guard, error) {
	fd, err := os.Open(f)
	if err != nil {
		return nil, nil, err
	}
	defer fd.Close()

	r := bufio.NewReader(fd)
	fp := &floorPlan{}
	var g *guard = nil
	line := 0
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, nil, err
		}
		lCopy := make([]byte, len(l))
		copy(lCopy, l)

		*fp = append(*fp, lCopy)
		if g == nil {
			for c := range l {
				if l[c] == 'v' {
					g = &guard{
						d: DOWN,
						p: position{c, line},
					}
				} else if l[c] == '^' {
					g = &guard{
						d: UP,
						p: position{c, line},
					}
				} else if l[c] == '<' {
					g = &guard{
						d: LEFT,
						p: position{c, line},
					}
				} else if l[c] == '>' {
					g = &guard{
						d: RIGHT,
						p: position{c, line},
					}
				}
			}
		}
		line++
	}
	return fp, g, nil
}

func main() {
	inputFlag := flag.String("input", "", "-input <input file>")

	flag.Parse()

	if inputFlag == nil {
		log.Fatal("empty input url")
	}
	fp, g, err := parse(*inputFlag)
	if err != nil {
		log.Fatal(err)
	}
	fp2, g2, _ := parse(*inputFlag)

	u := g.walk(fp)
	log.Printf("Unique positions: %d", u)

	obs := g2.walk2(fp2)
	log.Printf("Possible obstacles: %d", obs)
}
