package main

type Tree struct {
	value       int
	left, right *Tree
}

// Sorts values in place
func Sort(values []int) {
	var root *Tree

	for _, v := range values {
		root = add(root, v)
	}

	appendValues(values[:0], root)
}

// append the elements of t to values in order
// and return the resulting slice
func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *Tree, value int) *Tree {
	if t == nil {
		// Equivalent to return &Tree{value: value}
		t = new(Tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
