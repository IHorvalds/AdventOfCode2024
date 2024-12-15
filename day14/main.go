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
)

type position struct {
	x, y int
}

type velocity position

type robot struct {
	p position
	v velocity
}

var (
	width  = 101
	height = 103
)

func countRobots(rbts []robot) int {
	c00 := 0
	c01 := 0
	c10 := 0
	c11 := 0
	for _, r := range rbts {
		if r.p.x == (width / 2) {
			continue
		}

		if r.p.y == (height / 2) {
			continue
		}

		if r.p.x < width/2 && r.p.y < height/2 {
			c00++
		}
		if r.p.x > width/2 && r.p.y < height/2 {
			c01++
		}
		if r.p.x < width/2 && r.p.y > height/2 {
			c10++
		}
		if r.p.x > width/2 && r.p.y > height/2 {
			c11++
		}
	}
	return c00 * c01 * c10 * c11
}

func (r *robot) advance(seconds int) {
	for range seconds {
		r.p.x = (r.p.x + r.v.x) % width
		if r.p.x < 0 {
			r.p.x += width
		}
		r.p.y = (r.p.y + r.v.y) % height
		if r.p.y < 0 {
			r.p.y += height
		}
	}
}

func parseRobot(line string) (robot, error) {
	reg := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	ms := reg.FindStringSubmatch(line)
	if len(ms) != 5 {
		return robot{}, errors.New("invalid robot definition")
	}

	px, err := strconv.Atoi(ms[1])
	if err != nil {
		return robot{}, err
	}

	py, err := strconv.Atoi(ms[2])
	if err != nil {
		return robot{}, err
	}

	vx, err := strconv.Atoi(ms[3])
	if err != nil {
		return robot{}, err
	}

	vy, err := strconv.Atoi(ms[4])
	if err != nil {
		return robot{}, err
	}

	return robot{
		p: position{px, py},
		v: velocity{vx, vy},
	}, nil
}

func parse(f string) ([]robot, error) {
	fd, err := os.Open(f)
	if err != nil {
		return []robot{}, err
	}
	defer fd.Close()

	rbts := []robot{}
	r := bufio.NewReader(fd)
	for {
		l, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return []robot{}, err
		}
		rbt, err := parseRobot(l)
		if err != nil {
			return []robot{}, err
		}
		rbts = append(rbts, rbt)
	}

	return rbts, nil
}

func main() {
	inputFlag := flag.String("input", "", "-input <input file>")
	widthFlag := flag.Int("width", 101, "-width <# blocks>")
	heightFlag := flag.Int("height", 103, "-height <# blocks>")

	flag.Parse()

	if inputFlag == nil {
		log.Fatal("empty input url")
	}

	width = *widthFlag
	height = *heightFlag

	rbts, err := parse(*inputFlag)
	if err != nil {
		log.Fatal(err)
	}

	for i := range rbts {
		rbts[i].advance(100)
	}

	log.Printf("Safety factor: %d", countRobots(rbts))
}
