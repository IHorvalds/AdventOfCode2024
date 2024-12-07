package main

import "testing"

func TestSmallInput(t *testing.T) {
	fp, g, err := parse("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	u := g.walk(fp)
	if u != 41 {
		t.Fatalf("Expected 41 unique positions, got %d", u)
	}
}

func TestSmallInput2(t *testing.T) {
	fp, g, err := parse("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	u := g.walk2(fp)
	if u != 6 {
		t.Fatalf("Expected 6 possible obstacles, got %d", u)
	}
}
