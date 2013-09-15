package iter

import "fmt"

// ExampleIntIterable demonstrates an in-order iterator over a binary-search
// tree. This is the serial answer to the parallel tree-comparison example at
// http://golang.org/doc/play/tree.go.
func ExampleIntIterable() {
	var b *bst
	fmt.Printf("Do empty iterators have a next item? %t\n", b.Iter().HasNext())
	for _, value := range [...]int{3, 1, 5, 0, 6, 2, 4} {
		b = b.Insert(value)
	}
	fmt.Print("Elements: ")
	for it := b.Iter(); it.HasNext(); {
		// In this example, we can safely ignore the error.
		value, _ := it.Next()
		fmt.Printf("%d ", value)
	}
	// Output:
	// Do empty iterators have a next item? false
	// Elements: 0 1 2 3 4 5 6
}

// In-order iterator over numbers stored in a bst
type bstIterator struct {
	stack    []*bst // stack of visited positions
	currNode *bst
}

// Iter returns an in-order iterator over the elements of the tree.
func (t *bst) Iter() IntIterator {
	return &bstIterator{stack: make([]*bst, 0), currNode: t}
}

// HasNext returns true if calling Next safely returns the next item in the
// iterator.
func (it *bstIterator) HasNext() bool {
	if it == nil {
		return false
	}
	if len(it.stack) == 0 && it.currNode == nil {
		it.stack = nil // free memory
		return false
	}
	return true
}

// Next returns the next available item in the iterator if HasNext returns true,
// and EmptyIteratorError as the error value otherwise.
func (it *bstIterator) Next() (int, error) {
	// As a general matter, Next's code shouldn't repeat HasNext's. We're just
	// being a bit repetitive here to properly implement the recursive algorithm
	// with an explicit stack.
	for it.HasNext() {
		if it.currNode != nil {
			it.stack = append(it.stack, it.currNode) // Push
			it.currNode = it.currNode.left
		} else {
			it.currNode = it.stack[len(it.stack)-1] // Pop, part 1
			it.stack = it.stack[:len(it.stack)-1]   // Pop, part 2
			value := it.currNode.value
			it.currNode = it.currNode.right
			return value, nil
		}
	}
	return 0, EmptyIteratorError
}

// Unbalanced binary search tree.
type bst struct {
	value       int
	left, right *bst
}

// Insert a value into the subtree rooted at t
func (t *bst) Insert(value int) *bst {
	switch {
	case t == nil:
		return &bst{value: value}
	case value < t.value:
		t.left = t.left.Insert(value)
	case value > t.value:
		t.right = t.right.Insert(value)
	}
	return t
}

// Iter returns the receiver object.
func (it *bstIterator) Iter() IntIterator { return it }
