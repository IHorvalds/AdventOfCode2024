package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"log"
	"os"
)

type position struct {
	x, y int
}

type antennae map[byte][]position
type roofMap [][]byte

type positionSet map[position]struct{}

func placeAntinode(from, to position) position {
	return position{
		x: to.x + (to.x - from.x),
		y: to.y + (to.y - from.y),
	}
}

func placeAllAntinodes(from, to position, maxX, maxY int) positionSet {
	p := positionSet{}
	p[from] = struct{}{}
	p[to] = struct{}{}

	dx := to.x - from.x
	dy := to.y - from.y

	n := to
	for {
		n = position{
			x: n.x + dx,
			y: n.y + dy,
		}
		if n.x >= 0 && n.x < maxX && n.y >= 0 && n.y < maxY {
			p[n] = struct{}{}
		} else {
			break
		}
	}

	n = from
	for {
		n = position{
			x: n.x - dx,
			y: n.y - dy,
		}
		if n.x >= 0 && n.x < maxX && n.y >= 0 && n.y < maxY {
			p[n] = struct{}{}
		} else {
			break
		}
	}

	return p
}

func findAntinodes(ant []position, maxX, maxY int) positionSet {
	p := positionSet{}
	for i := range ant {
		for j := i + 1; j < len(ant); j++ {
			an := placeAntinode(ant[i], ant[j])
			if an.x >= 0 && an.x < maxX && an.y >= 0 && an.y < maxY {
				p[an] = struct{}{}
			}

			an = placeAntinode(ant[j], ant[i])
			if an.x >= 0 && an.x < maxX && an.y >= 0 && an.y < maxY {
				p[an] = struct{}{}
			}
		}
	}
	return p
}

func findAntinodes2(ant []position, maxX, maxY int) positionSet {
	p := positionSet{}
	for i := range ant {
		for j := i + 1; j < len(ant); j++ {
			for an := range placeAllAntinodes(ant[i], ant[j], maxX, maxY) {
				p[an] = struct{}{}
			}
		}
	}
	return p
}

func (a *antennae) insert(freq byte, p position) {
	if _, ok := (*a)[freq]; !ok {
		(*a)[freq] = []position{p}
	} else {
		(*a)[freq] = append((*a)[freq], p)
	}
}

func merge(a, b positionSet) positionSet {
	p := positionSet{}
	for pos := range a {
		p[pos] = struct{}{}
	}
	for pos := range b {
		p[pos] = struct{}{}
	}
	return p
}

func (roof *roofMap) findAllAntinodes(a antennae) int {
	p := positionSet{}
	for _, ant := range a {
		p = merge(p, findAntinodes(ant, len((*roof)[0]), len(*roof)))
	}
	return len(p)
}

func (roof *roofMap) findAllAntinodes2(a antennae) int {
	p := positionSet{}
	for _, ant := range a {
		p = merge(p, findAntinodes2(ant, len((*roof)[0]), len(*roof)))
	}
	return len(p)
}

func parse(f string) (roofMap, antennae, error) {
	fd, err := os.Open(f)
	if err != nil {
		return roofMap{}, antennae{}, err
	}
	defer fd.Close()

	roof := roofMap{}
	ant := antennae{}
	line := 0
	r := bufio.NewReader(fd)
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return roofMap{}, antennae{}, err
		}

		lCopy := make([]byte, len(l))
		copy(lCopy, l)

		roof = append(roof, lCopy)

		for i := range lCopy {
			if lCopy[i] != '.' {
				ant.insert(lCopy[i], position{
					x: i,
					y: line,
				})
			}
		}
		line++
	}
	return roof, ant, nil
}

func main() {
	inputFlag := flag.String("input", "", "-input <input file>")

	flag.Parse()

	if inputFlag == nil {
		log.Fatal("empty input url")
	}

	roof, ant, err := parse(*inputFlag)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Antinodes: %d", roof.findAllAntinodes(ant))
	log.Printf("Antinodes 2: %d", roof.findAllAntinodes2(ant))
}
