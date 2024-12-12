package main

import (
	"flag"
	"log"
	"os"
	"slices"
	"strconv"
)

type block struct {
	id  int
	pos int
}

type disk struct {
	compactRepr     []int
	head, tail, end uint64
	tailIdx         int
}

type checksum uint64

func (c *checksum) update(c2 checksum) {
	*c += c2
}

func calcChecksum(bs ...block) checksum {
	c := checksum(0)
	for _, b := range bs {
		if b.id != -1 {
			c.update(checksum(b.id * b.pos))
		}
	}
	return c
}

func (d disk) getSequenceStart() int {
	s := 0
	for i := len(d.compactRepr) - 1; i >= d.tailIdx; i-- {
		s += d.compactRepr[i]
	}
	return s
}

func (d *disk) getLast(count int) []int {
	ls := make([]int, count)
	c := 0
	for d.tailIdx > 0 && c < count {
		for ; d.tail > d.end-uint64(d.getSequenceStart()) && c < count; d.tail-- {
			if d.tailIdx%2 == 0 {
				ls[c] = d.tailIdx / 2
				c++
			}
		}
		if c < count {
			d.tailIdx--
		}
	}
	return ls
}

func part1(d disk) checksum {
	c := checksum(0)
	for i := range d.compactRepr {
		if i%2 == 0 { // file
			for range d.compactRepr[i] {
				if d.head >= d.tail {
					break
				}
				c += calcChecksum(block{
					id:  i / 2,
					pos: int(d.head),
				})
				d.head++
			}
		} else { // free space
			for _, v := range d.getLast(d.compactRepr[i]) {
				if d.head >= d.tail {
					break
				}
				c += calcChecksum(block{
					id:  v,
					pos: int(d.head),
				})
				d.head++
			}
		}

		if d.head >= d.tail {
			break
		}
	}
	return c
}

func part2(d disk) checksum {
	c := checksum(0)
	seen := []int{}
	for h := 0; h < len(d.compactRepr); h++ {
		if h%2 == 0 {
			if !slices.Contains(seen, h) {
				for range d.compactRepr[h] {
					c += calcChecksum(block{
						id:  h / 2,
						pos: int(d.head),
					})
					d.head++
				}
			} else {
				d.head += uint64(d.compactRepr[h])
			}
		} else {
			spaces := d.compactRepr[h]

			for t := len(d.compactRepr) - 1; t > h && spaces > 0; {
				if t%2 != 0 {
					t--
					continue
				}
				if d.compactRepr[t] <= spaces && !slices.Contains(seen, t) {
					seen = append(seen, t)
					for range d.compactRepr[t] {
						c += calcChecksum(block{
							id:  t / 2,
							pos: int(d.head),
						})
						d.head++
					}

					spaces -= d.compactRepr[t]
				}

				t -= 2
			}
			d.head += uint64(spaces)
		}

	}

	return c
}

func parse(f string) (disk, error) {
	l, err := os.ReadFile(f)
	if err != nil {
		return disk{}, err
	}
	is := make([]int, len(l))
	var t uint64 = 0
	for i, b := range l {
		is[i], err = strconv.Atoi(string(b))
		if err != nil {
			return disk{}, err
		}
		t += uint64(is[i])
	}
	d := disk{
		compactRepr: is,
		head:        0,
		tail:        t,
		end:         t,
		tailIdx:     len(is) - 1,
	}

	return d, nil
}

func main() {
	inputFlag := flag.String("input", "", "-input <input file>")
	part := flag.String("part", "", "-part [1 or 2]")

	flag.Parse()

	if inputFlag == nil {
		log.Fatal("empty input url")
	}

	d, err := parse(*inputFlag)
	if err != nil {
		log.Fatal(err)
	}

	if *part == "1" {
		log.Printf("Checksum 1: %d", part1(d))
	}
	if *part == "2" {
		log.Printf("Checksum 2: %d", part2(d))
	}
}
