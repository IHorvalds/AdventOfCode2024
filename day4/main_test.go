package main

import (
	"testing"
)

func TestInput(t *testing.T) {
	v, err := parse("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	c := part1(&v)

	if c != 18 {
		t.Fatalf("Expected 18, got %d", c)
	}
}

func TestInputPart2(t *testing.T) {
	v, err := parse("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	c := part2(&v)

	if c != 9 {
		t.Fatalf("Expected 9, got %d", c)
	}
}
