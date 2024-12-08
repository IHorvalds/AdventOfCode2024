package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	res      uint64
	operands []uint64
}

type OP int

const (
	ADD OP = iota
	MUL
	CON
)

func operate(a, b uint64, o OP) uint64 {
	switch o {
	case ADD:
		return a + b
	case MUL:
		return a * b
	case CON:
		return uint64(math.Pow10(int(math.Floor(math.Log10(float64(b))))+1))*a + b
	}
	panic(fmt.Sprintf("Unknown operator %d", o))
}

func apply(operands []uint64, ops []OP) uint64 {
	if len(ops) != len(operands)-1 {
		panic(fmt.Sprintf("invalid number of operands and operators: %d and %d", len(ops), len(operands)))
	}

	t := operands[0]
	for o := range ops {
		u := operands[o+1]
		t = operate(t, u, ops[o])
	}
	return t
}

// sudoku style...
func incrementOps(ops []OP) bool {

	if len(ops) == 0 {
		return false
	}

	if ops[0] < CON {
		ops[0]++
		return true
	} else {
		ops[0] = ADD
	}

	return incrementOps(ops[1:])
}

func findOperators(eq *equation) []OP {
	ops := make([]OP, len((*eq).operands)-1)
	for {
		if eq.res == apply(eq.operands, ops) {
			return ops
		}
		if !incrementOps(ops) {
			break
		}
	}

	return nil
}

func parse(f string) ([]equation, error) {
	fd, err := os.Open(f)
	if err != nil {
		return []equation{}, err
	}
	defer fd.Close()

	r := bufio.NewReader(fd)
	e := []equation{}
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return []equation{}, err
		}
		s := string(l)
		ns := strings.Split(s, ": ")
		if len(ns) != 2 {
			panic("invalid input")
		}

		res, err := strconv.ParseUint(ns[0], 10, 64)
		if err != nil {
			return []equation{}, err
		}

		ops := []uint64{}
		for _, x := range strings.Split(ns[1], " ") {
			o, err := strconv.ParseUint(x, 10, 64)
			if err != nil {
				return []equation{}, err
			}
			ops = append(ops, o)
		}
		e = append(e, equation{
			res:      res,
			operands: ops,
		})
	}

	return e, nil
}

func main() {
	inputFlag := flag.String("input", "", "-input <input file>")

	flag.Parse()

	if inputFlag == nil {
		log.Fatal("empty input url")
	}

	eqs, err := parse(*inputFlag)
	if err != nil {
		log.Fatal(err)
	}

	var c uint64 = 0
	for i := range eqs {
		ops := findOperators(&eqs[i])
		if ops != nil {
			c += eqs[i].res
		}
	}

	log.Printf("Total %d", c)
}
