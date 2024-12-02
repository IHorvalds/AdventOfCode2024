package main

import (
	"log"
	"testing"
)

func TestSafety(t *testing.T) {
	// 7 6 4 2 1 -> Safe
	// 1 2 7 8 9 -> Unsafe (delta = +5)
	// 9 7 6 2 1 -> Unsafe (delta = -4)
	// 1 3 2 4 5 -> Unsafe (delta changed sign at 3->2)
	// 8 6 4 4 1 -> Unsafe (delta == 0)
	// 1 3 6 7 9 -> Safe
	in := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	res := []safety{SAFE, UNSAFE, UNSAFE, UNSAFE, UNSAFE, SAFE}
	for i := range in {
		s, _ := isSafe(&in[i])
		if s != res[i] {
			t.Fail()
			log.Printf("For %d, got %d, expected %d", i, s, res[i])
		}
	}
}

func TestSafeish(t *testing.T) {
	in := []struct {
		inputs []int
		safe   safety
	}{
		{[]int{7, 6, 4, 2, 1}, SAFE},
		{[]int{1, 2, 7, 8, 9}, UNSAFE},
		{[]int{9, 7, 6, 2, 1}, UNSAFE},
		{[]int{1, 3, 2, 4, 5}, SAFE},
		{[]int{8, 6, 4, 4, 1}, SAFE},
		{[]int{1, 3, 6, 7, 9}, SAFE},
		{[]int{5, 4, 3, 1, 2}, SAFE},
		{[]int{7, 1, 2, 3, 4}, SAFE},
		{[]int{7, 1, 2, 3, 3}, UNSAFE},
		{[]int{1, 7, 3, 4, 5}, SAFE},
		{[]int{1, 3, 4, 5, 9}, SAFE},
	}
	for i := range in {
		s := isSafeish(&in[i].inputs)
		if s != in[i].safe {
			t.Fail()
			log.Printf("For %d, got %d, expected %d", i, s, in[i].safe)
		}
	}
}
