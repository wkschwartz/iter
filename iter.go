/*
Package iter defines iterable and iterator interfaces -- protocols really.

Iterators are clunkier and likely slower than for-range statements, so they
aren't for iterating over slices, maps, or channels. Rather, they're most useful
in contexts where go does not provide an obvious way to iterate.

This package provides examples of the interface for elements of type interface{},
int, and string.
*/
package iter

import "errors"

// EmptyIteratorError is optionally the error return value for an iterator's
// Next method when its HasNext method returns false because the iterator is
// empty.
var EmptyIteratorError error = errors.New("iter: empty iterator")

// Iterable containers allow iteration over their contents with incantations
// like
//     for it := container.Iter(); it.HasNext(); {
//         item, err := it.Next()
//         // operations with item
//     }
// Not all iterators will use the error return value, and you can safely ignore
// the error return value in those cases. Others iterators, such as ones that
// iterate over files, may use them to indicate problems in the underlying
// data. Make sure HasNext returns true before calling Next.
type Iterable interface {
	// Iter returns a new Iterator instance starting from a position independant
	// of the state of any other Iterators.
	Iter() Iterator
}

// Iterators are made with an Iterable's Iter method. Then the caller calls Next
// as long as HasNext returns true. The behavior of an iterator is undefined
// when the underlying container is modified after Iter returns the iterator
// (except that calling an iterator's Iter method returns the receiver
// instance).
type Iterator interface {
	// Iter returns the receiver.
	Iterable
	// HasNext returns false if the iterator is empty, guaranteeing that Next's
	// error return value is not EmptyIteratorError.
	HasNext() bool
	// Next returns the next item in the iterator if HasNext returns true and
	// the return error value is nil. If HasNext returns false, Next may return
	// EmptyIteratorError as the error, but the main return value is undefined.
	Next() (interface{}, error)
}

// Iteratable over int elements.
type IntIterable interface {
	Iter() IntIterator
}

type IntIterator interface {
	IntIterable
	HasNext() bool
	Next() (int, error)
}

// Iteratable over string elements.
type StringIterable interface {
	Iter() StringIterator
}

type StringIterator interface {
	StringIterable
	HasNext() bool
	Next() (string, error)
}
