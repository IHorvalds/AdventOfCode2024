package main

import (
	"log"
	"testing"
)

func TestSplit(t *testing.T) {
	if a, b := splitDigits(1234); a != 12 || b != 34 {
		log.Printf("1234 was split into %d %d", a, b)
		t.Fail()
	}

	if a, b := splitDigits(9999); a != 99 || b != 99 {
		log.Printf("9999 was split into %d %d", a, b)
		t.Fail()
	}

	if a, b := splitDigits(90); a != 9 || b != 0 {
		log.Printf("90 was split into %d %d", a, b)
		t.Fail()
	}

	if a, b := splitDigits(9900); a != 99 || b != 0 {
		log.Printf("9900 was split into %d %d", a, b)
		t.Fail()
	}

	if a, b := splitDigits(4000); a != 40 || b != 00 {
		log.Printf("4000 was split into %d %d", a, b)
		t.Fail()
	}
}
