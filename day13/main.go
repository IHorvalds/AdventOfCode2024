package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type button struct {
	x, y int // increments along the x and y axes, respectively
}

type prize struct {
	x, y int // coordinates of the prize
}

type game struct {
	a, b button
	p    prize
}

func (g game) intersection() (int, int) {
	// calculating the intersection of
	// 0 = a.x * C1 + b.x * C2 - p.x
	// and
	// 0 = a.y * C1 + b.y * C2 - p.y
	x := (-g.p.y*g.b.x + g.p.x*g.b.y) / (g.a.x*g.b.y - g.a.y*g.b.x) // presses for button A
	y := (g.p.x - g.a.x*x) / g.b.x                                  // presses for button B
	return x, y
}

func checkIntersection(g *game, x, y int) bool {
	return g.p.x == x*g.a.x+y*g.b.x && g.p.y == x*g.a.y+y*g.b.y
}

func parseButton(s string) (button, error) {
	r := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
	ms := r.FindStringSubmatch(s)
	if len(ms) != 3 {
		return button{}, errors.New("invalid button format")
	}
	x, err := strconv.Atoi(ms[1])
	if err != nil {
		return button{}, err
	}

	y, err := strconv.Atoi(ms[2])
	if err != nil {
		return button{}, err
	}

	return button{x, y}, nil
}

func parsePrize(s string) (prize, error) {
	r := regexp.MustCompile(`X=(\d+), Y=(\d+)`)
	ms := r.FindStringSubmatch(s)
	if len(ms) != 3 {
		return prize{}, errors.New("invalid prize format")
	}
	x, err := strconv.Atoi(ms[1])
	if err != nil {
		return prize{}, err
	}

	y, err := strconv.Atoi(ms[2])
	if err != nil {
		return prize{}, err
	}

	return prize{x, y}, nil
}

func parseGame(r *bufio.Reader) (game, error) {
	g := game{}

	l, err := r.ReadString('\n')
	if err != nil {
		if errors.Is(err, io.EOF) {
			err = errors.New("unexpected EOF")
		}
		return game{}, err
	}
	if !strings.HasPrefix(l, "Button A: ") {
		return game{}, errors.New("missing button A")
	}
	if b, err := parseButton(l); err != nil {
		return game{}, err
	} else {
		g.a = b
	}

	l, err = r.ReadString('\n')
	if err != nil {
		if errors.Is(err, io.EOF) {
			err = errors.New("unexpected EOF")
		}
		return game{}, err
	}
	if !strings.HasPrefix(l, "Button B: ") {
		return game{}, errors.New("missing button B")
	}
	if b, err := parseButton(l); err != nil {
		return game{}, err
	} else {
		g.b = b
	}

	l, err = r.ReadString('\n')
	if err != nil {
		if errors.Is(err, io.EOF) {
			err = errors.New("unexpected EOF")
		}
		return game{}, err
	}
	if !strings.HasPrefix(l, "Prize: ") {
		return game{}, errors.New("missing prize")
	}
	if p, err := parsePrize(l); err != nil {
		return game{}, err
	} else {
		g.p = p
	}

	l, err = r.ReadString('\n')
	if err != nil {
		if !errors.Is(err, io.EOF) {
			return game{}, err
		}

		if l != "" {
			return game{}, errors.New("unexpected non-empty new line")
		}
	}
	return g, err
}

func parse(f string) ([]game, error) {
	fd, err := os.Open(f)
	if err != nil {
		return []game{}, err
	}
	defer fd.Close()

	r := bufio.NewReader(fd)
	g := []game{}
	for {
		if gm, err := parseGame(r); err == nil {
			g = append(g, gm)
		} else {
			if errors.Is(err, io.EOF) {
				g = append(g, gm)
				break
			}
			return []game{}, err
		}
	}
	return g, nil
}

func main() {
	inputFlag := flag.String("input", "", "-input <input file>")

	flag.Parse()

	if inputFlag == nil {
		log.Fatal("empty input url")
	}

	gms, err := parse(*inputFlag)
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	for _, g := range gms {
		a, b := g.intersection()
		if checkIntersection(&g, a, b) {
			total += a*3 + b
		}
	}

	log.Printf("Minimum total tokens: %d", total)
}
