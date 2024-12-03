package main

import (
	"log"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	s := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	n := parseBytes(s, false)

	if n != 161 {
		t.Fatalf("Expected 161, got %d", n)
	}
}

func TestPart2(t *testing.T) {
	s := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	n := parseBytes(s, true)

	if n != 48 {
		t.Fatalf("Expected 161, got %d", n)
	}
}

func TestPart1File(t *testing.T) {
	bs, err := os.ReadFile("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	parseBytes(string(bs), false)
}
