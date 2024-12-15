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

func isCorner(plant byte, n1, n2 *byte) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if n1 == nil && n2 != nil && *n2 != plant {
		return true
	}

	if n2 == nil && n1 != nil && *n1 != plant {
		return true
	}

	if n1 != nil && n2 != nil && *n1 != plant && *n2 != plant {
		return true
	}

	return false
}

func countCorners(g garden, pos position) int {
	c := 0
	var n *byte = nil
	var s *byte = nil
	var e *byte = nil
	var w *byte = nil
	if pos.x > 0 {
		n = new(byte)
		*n = g[pos.x-1][pos.y]
	}

	if pos.x < len(g)-1 {
		s = new(byte)
		*s = g[pos.x+1][pos.y]
	}

	if pos.y > 0 {
		w = new(byte)
		*w = g[pos.x][pos.y-1]
	}

	if pos.y < len(g[0])-1 {
		e = new(byte)
		*e = g[pos.x][pos.y+1]
	}

	plant := g[pos.x][pos.y]

	// convex corners
	if isCorner(plant, w, n) {
		c++
	}
	if isCorner(plant, n, e) {
		c++
	}
	if isCorner(plant, e, s) {
		c++
	}
	if isCorner(plant, s, w) {
		c++
	}

	// concave corners
	if w != nil && n != nil && *w == plant && *n == plant && g[pos.x-1][pos.y-1] != plant {
		c++
	}
	if n != nil && e != nil && *n == plant && *e == plant && g[pos.x-1][pos.y+1] != plant {
		c++
	}
	if e != nil && s != nil && *e == plant && *s == plant && g[pos.x+1][pos.y+1] != plant {
		c++
	}
	if s != nil && w != nil && *s == plant && *w == plant && g[pos.x+1][pos.y-1] != plant {
		c++
	}

	return c
}

func areaAndCorners(g garden, vMap [][]bool, pos position) (int, int) {
	corners := 0
	area := 0
	plant := g[pos.x][pos.y]

	vst := list.New().Init()
	vst.PushBack(pos)
	for vst.Len() > 0 {
		p := vst.Remove(vst.Front()).(position)
		if vMap[p.x][p.y] {
			continue
		}
		area++
		vMap[p.x][p.y] = true
		corners += countCorners(g, p)

		if p.x > 0 {
			left := g[p.x-1][p.y]
			if left == plant && !vMap[p.x-1][p.y] {
				vst.PushBack(position{p.x - 1, p.y})
			}
		}

		if p.x < len(g)-1 {
			left := g[p.x+1][p.y]
			if left == plant && !vMap[p.x+1][p.y] {
				vst.PushBack(position{p.x + 1, p.y})
			}
		}

		if p.y > 0 {
			left := g[p.x][p.y-1]
			if left == plant && !vMap[p.x][p.y-1] {
				vst.PushBack(position{p.x, p.y - 1})
			}
		}

		if p.y < len(g[0])-1 {
			left := g[p.x][p.y+1]
			if left == plant && !vMap[p.x][p.y+1] {
				vst.PushBack(position{p.x, p.y + 1})
			}
		}
	}

	return area, corners
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

func price2(g garden, v [][]bool) int {
	pr := 0
	for i := range len(g) {
		for j := range len(g[i]) {
			if !v[i][j] {
				a, p := areaAndCorners(g, v, position{i, j})
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
	discountFlag := flag.Bool("discount", false, "-discount")

	flag.Parse()

	if inputFlag == nil {
		log.Fatal("empty input url")
	}

	g, v, err := parse(*inputFlag)
	if err != nil {
		log.Fatal(err)
	}

	if !*discountFlag {
		log.Printf("Fence price: %d", price(g, v))
	} else {
		log.Printf("Fence price discounted: %d", price2(g, v))
	}
}
