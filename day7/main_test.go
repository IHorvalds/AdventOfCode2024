package main

import (
	"log"
	"slices"
	"testing"
)

func TestIncrement(t *testing.T) {
	v := [][]OP{
		{ADD, ADD, ADD, ADD},
		{MUL, ADD, ADD, ADD},
		{CON, ADD, ADD, ADD},
		{ADD, MUL, ADD, ADD},
		{MUL, MUL, ADD, ADD},
		{CON, MUL, ADD, ADD},
		{ADD, CON, ADD, ADD},
		{MUL, CON, CON, CON},
		{CON, CON, CON, CON},
	}
	if !incrementOps(v[0]) || slices.Compare(v[0], v[1]) != 0 {
		log.Print("Failed for increment 1")
		t.Fail()
	}
	if !incrementOps(v[1]) || slices.Compare(v[1], v[2]) != 0 {
		log.Print("Failed for increment 2")
		t.Fail()
	}
	if !incrementOps(v[2]) || slices.Compare(v[2], v[3]) != 0 {
		log.Print("Failed for increment 3")
		t.Fail()
	}
	if !incrementOps(v[3]) || slices.Compare(v[3], v[4]) != 0 {
		log.Print("Failed for increment 4")
		t.Fail()
	}
	if !incrementOps(v[4]) || slices.Compare(v[4], v[5]) != 0 {
		log.Print("Failed for increment 5")
		t.Fail()
	}
	if !incrementOps(v[5]) || slices.Compare(v[5], v[6]) != 0 {
		log.Print("Failed for increment 6")
		t.Fail()
	}
	if !incrementOps(v[7]) || slices.Compare(v[7], v[8]) != 0 {
		log.Print("Failed for increment 7")
		t.Fail()
	}
	if incrementOps(v[8]) {
		log.Print("Failed for increment 8")
		t.Fail()
	}
}

func TestNumConcat(t *testing.T) {
	x := operate(12, 34, CON)
	if x != 1234 {
		t.Fatalf("Got %d", x)
	}
}
