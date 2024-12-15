package main

import "testing"

func TestLineIntersection(t *testing.T) {
	g := game{
		a: button{94, 34},
		b: button{22, 67},
		p: prize{8400, 5400},
	}

	if a, b := g.intersection(); a != 80 || b != 40 {
		t.Fatalf("Expected 80, 40, got: %d, %d", a, b)
	}

	g = game{
		a: button{17, 86},
		b: button{84, 37},
		p: prize{7870, 6450},
	}

	if a, b := g.intersection(); a != 38 || b != 86 {
		t.Fatalf("Expected 38, 86, got: %d, %d", a, b)
	}
}

func TestLineNoIntersection(t *testing.T) {
	g := game{
		a: button{26, 66},
		b: button{67, 21},
		p: prize{12748, 12176},
	}

	if a, b := g.intersection(); checkIntersection(&g, a, b) {
		t.Fatalf("Expected no solution, got: %d, %d", a, b)
	}
}
