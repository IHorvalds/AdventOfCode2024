package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"time"
)

type position struct {
	x, y int
}

type velocity position

type robot struct {
	p position
	v velocity
}

type robots []robot

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

func (rbts robots) display(second int) {

	m := map[position]int{}
	for _, r := range rbts {
		if _, ok := m[r.p]; !ok {
			m[r.p] = 1
		} else {
			m[r.p]++
		}
	}
	fmt.Print("\033[H\033[2J")
	fmt.Printf("Second: %d\n\n", second)

	for i := range height {
		for j := range width {
			p := position{j, i}
			if c, ok := m[p]; ok {
				fmt.Printf("%d", c)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Scan("Next?")
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

func parse(f string) (robots, error) {
	fd, err := os.Open(f)
	if err != nil {
		return robots{}, err
	}
	defer fd.Close()

	rbts := robots{}
	r := bufio.NewReader(fd)
	for {
		l, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return robots{}, err
		}
		rbt, err := parseRobot(l)
		if err != nil {
			return robots{}, err
		}
		rbts = append(rbts, rbt)
	}

	return rbts, nil
}

func main() {
	inputFlag := flag.String("input", "", "-input <input file>")
	widthFlag := flag.Int("width", 101, "-width <# blocks>")
	heightFlag := flag.Int("height", 103, "-height <# blocks>")
	watchOutput := flag.Bool("watch", false, "-watch")
	secondFlag := flag.Int("second", -1, "-second # The second to display")

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

	if !*watchOutput {
		if *secondFlag != -1 {
			for i := range rbts {
				rbts[i].advance(100)
			}

			log.Printf("Safety factor: %d", countRobots(rbts))
		} else {
			for i := range rbts {
				rbts[i].advance(*secondFlag)
			}

			rbts.display(*secondFlag)
		}
	} else {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		running := true

		c2 := make(chan struct{})
		go func() {
			s := 0
			for running {
				for i := range rbts {
					rbts[i].advance(1)
				}

				rbts.display(s)
				s++
				time.Sleep(time.Millisecond * 165)
			}

			c2 <- struct{}{}
		}()

		<-c
		running = false

		<-c2
	}
}
