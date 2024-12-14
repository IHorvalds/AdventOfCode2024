package main

import (
	"bufio"
	"container/list"
	"errors"
	"flag"
	"io"
	"log"
	"os"
)

type garden [][]byte

type position struct {
	x, y int
}

func areaAndPerimeter(g garden, vMap [][]bool, pos position) (int, int) {
	perimeter := 0
	area := 0

	vst := list.New().Init()
	plant := g[pos.x][pos.y]
	vst.PushBack(pos)
	for vst.Len() > 0 {
		p := vst.Remove(vst.Front()).(position)
		if vMap[p.x][p.y] {
			continue
		}
		area++
		vMap[p.x][p.y] = true
		if p.x > 0 {
			left := g[p.x-1][p.y]
			if left != plant {
				perimeter++
			} else if !vMap[p.x-1][p.y] {
				vst.PushBack(position{p.x - 1, p.y})
			}
		} else {
			perimeter++
		}

		if p.x < len(g)-1 {
			left := g[p.x+1][p.y]
			if left != plant {
				perimeter++
			} else if !vMap[p.x+1][p.y] {
				vst.PushBack(position{p.x + 1, p.y})
			}
		} else {
			perimeter++
		}

		if p.y > 0 {
			left := g[p.x][p.y-1]
			if left != plant {
				perimeter++
			} else if !vMap[p.x][p.y-1] {
				vst.PushBack(position{p.x, p.y - 1})
			}
		} else {
			perimeter++
		}

		if p.y < len(g[0])-1 {
			left := g[p.x][p.y+1]
			if left != plant {
				perimeter++
			} else if !vMap[p.x][p.y+1] {
				vst.PushBack(position{p.x, p.y + 1})
			}
		} else {
			perimeter++
		}
	}

	return area, perimeter
}

func price(g garden, v [][]bool) int {
	pr := 0
	for i := range len(g) {
		for j := range len(g[i]) {
			if !v[i][j] {
				a, p := areaAndPerimeter(g, v, position{i, j})
				pr += a * p
			}
		}
	}
	return pr
}

func parse(f string) (garden, [][]bool, error) {
	fd, err := os.Open(f)
	if err != nil {
		return garden{}, [][]bool{}, err
	}
	defer fd.Close()

	g := garden{}
	v := [][]bool{}
	r := bufio.NewReader(fd)
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return garden{}, [][]bool{}, err
		}

		lCopy := make([]byte, len(l))
		copy(lCopy, l)

		v = append(v, make([]bool, len(l)))

		g = append(g, lCopy)
	}
	return g, v, nil
}

func main() {
	inputFlag := flag.String("input", "", "-input <input file>")

	flag.Parse()

	if inputFlag == nil {
		log.Fatal("empty input url")
	}

	g, v, err := parse(*inputFlag)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Fence price: %d", price(g, v))
}
