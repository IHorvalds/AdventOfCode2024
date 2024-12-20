package main

import (
	"container/list"
	"fmt"
	"iter"
)

type node struct {
	v     int
	vs    []int // v may appear many times in the input. We want to capture them all
	left  *node
	right *node
	bal   int
}

type AVL struct {
	root  *node
	count int
}

func rotateLeft(n *node) node {
	nCopy := *n
	right := n.right
	rleft := right.left

	right.left = &nCopy
	nCopy.right = rleft

	// newN := *(*n).right
	// if newN.left != nil {
	// 	*(nCopy.right) = *(newN.left)
	// 	*(newN.left) = nCopy
	// } else {
	// 	newN.left = &nCopy
	// }
	setBalance(right.left)
	setBalance(right.right)
	setBalance(right)
	return *right
}

func rotateRight(n *node) node {
	nCopy := *n
	left := n.left
	lright := left.right

	left.right = &nCopy
	nCopy.left = lright

	// if newN.right != nil {
	// 	*(nCopy.left) = *(newN.right)
	// 	*(newN.right) = nCopy
	// } else {
	// 	newN.right = &nCopy
	// }

	setBalance(left.left)
	setBalance(left.right)
	setBalance(left)
	return *left
}

func (t *AVL) insert(val int) {
	if t.root == nil {
		t.root = createNode(val)
		t.count = 1
	} else {
		insert(t, t.root, val)
	}
}

func insert(t *AVL, n *node, val int) {
	if val == n.v {
		n.vs = append(n.vs, val)
		t.count = t.count + 1
		return
	}

	if val < n.v {
		if n.left == nil {
			n.left = createNode(val)
			t.count = t.count + 1
		} else {
			insert(t, n.left, val)
		}
		n.bal = n.bal - 1
	} else {
		if n.right == nil {
			n.right = createNode(val)
			t.count = t.count + 1
		} else {
			insert(t, n.right, val)
		}
		n.bal = n.bal + 1
	}

	assertBalance(n.bal)

	// rebalance
	if n.bal >= 2 { // right heavy
		if n.right.bal > 0 {
			*n = rotateLeft(n)
		} else if n.right.bal < 0 {
			*(n.right) = rotateRight(n.right)
			*n = rotateLeft(n)
		}
	} else if n.bal <= -2 { // left heavy
		if n.left.bal < 0 {
			*n = rotateRight(n)
		} else if n.left.bal > 0 {
			*(n.left) = rotateLeft(n.left)
			*n = rotateRight(n)
		}
	}
	setBalance(n)
}

func createNode(v int) *node {
	return &node{
		v:     v,
		vs:    []int{v},
		left:  nil,
		right: nil,
		bal:   0,
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func setBalance(n *node) {
	if n == nil {
		return
	}
	b := 0
	if n.left != nil {
		b = b - 1 - abs(n.left.bal)
	}
	if n.right != nil {
		b = b + 1 + abs(n.right.bal)
	}
	n.bal = b
}

func addToList(n *node, l *list.List) {
	if n.left != nil {
		addToList(n.left, l)
	}

	for _, v := range n.vs {
		l.PushBack(v)
	}

	if n.right != nil {
		addToList(n.right, l)
	}
}

func (t *AVL) toList() *list.List {
	l := list.New()
	l.Init()

	if t.root != nil {
		addToList(t.root, l)
	}

	return l
}

func (n *node) yieldVal(yield func(int) bool) bool {
	if n.left != nil {
		if !n.left.yieldVal(yield) {
			return false
		}
	}
	for _, v := range n.vs {
		if !yield(v) {
			return false
		}
	}
	if n.right != nil {
		if !n.right.yieldVal(yield) {
			return false
		}
	}
	return true
}

func (a *AVL) traverse() iter.Seq[int] {
	return func(yield func(int) bool) {
		a.root.yieldVal(yield)
	}
}

func diff(a1 *AVL, a2 *AVL) int {
	diff := 0

	n1, s1 := iter.Pull(a1.traverse())
	n2, s2 := iter.Pull(a2.traverse())

	for {
		v1, ok := n1()
		if !ok {
			s2()
			break
		}
		v2, ok := n2()
		if !ok {
			s1()
			break
		}
		d := v1 - v2
		if d < 0 {
			d = -d
		}
		diff = diff + d
	}

	return diff
}

func assertBalance(b int) {
	if b > 2 || b < -2 {
		panic(fmt.Sprintf("Wrong balance %d", b))
	}
}
