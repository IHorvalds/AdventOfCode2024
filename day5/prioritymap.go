package main

import (
	"slices"
	"sort"
)

// the elements of the slice have higher priority than the associated key
type prioritySlice map[int][]int

func (p *prioritySlice) insert(k, v int) {
	if _, ok := (*p)[k]; !ok {
		(*p)[k] = []int{v}
	} else {
		(*p)[k] = append((*p)[k], v)
	}
}

func (p *prioritySlice) sortPriorities() {
	for i := range *p {
		slices.Sort((*p)[i])
	}
}

func (p prioritySlice) isBefore(i, j int) bool {
	v, ok := p[j]
	if !ok {
		return false
	}

	_, found := slices.BinarySearch(v, i)
	return found
}

func (p prioritySlice) checkList(v *[]int) bool {
	if len(*v) <= 1 {
		return true
	}

	prev := (*v)[0]
	for i := 1; i < len(*v); i++ {
		if !p.isBefore(prev, (*v)[i]) {
			return false
		}
		prev = (*v)[i]
	}

	return true
}

func (p prioritySlice) sortList(v *[]int) {
	sort.Slice(*v, func(i, j int) bool {
		return p.isBefore((*v)[i], (*v)[j])
	})
}
