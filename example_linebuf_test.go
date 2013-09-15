package iter

import (
	"bufio"
	"fmt"
	"strings"
)

// ExampleStringIterable demonstrates reading a file with an iterator. This is
// sort of a trivial example because the bufio.scanner already has a fairly
// simple interface for line-buffered reading. The point here is mainly to show
// the utility of having the error return value from Next.
func ExampleStringIterator() {
	scanner := bufio.NewScanner(strings.NewReader(jabberwocky))
	// Iterators are also iterable.
	var iterable StringIterable = &lineIterator{
		scanner: scanner, hasNext: scanner.Scan()}
	for it := iterable.Iter(); it.HasNext(); {
		line, err := it.Next()
		if err != nil {
			panic(err)
		}
		fmt.Println(line)
	}
	// Output:
	//
	// He took his vorpal sword in hand;
	//   Long time the manxome foe he sought—
	// So rested he by the Tumtum tree,
	//   And stood awhile in thought.
}

// lineIterator allows iterating over an io.Reader by line; it satisfies the
// Iterator interface.
type lineIterator struct {
	scanner *bufio.Scanner
	hasNext bool
}

// Iter returns the receiver.
func (it *lineIterator) Iter() StringIterator { return it }

// HasNext returns whether the input stream has another line.
func (it *lineIterator) HasNext() bool { return it.hasNext }

// Next returns the next line in the input stream, and any read error if one
// occurred. The ending newline is stripped.
func (it *lineIterator) Next() (s string, e error) {
	s = it.scanner.Text()
	e = it.scanner.Err()
	if !it.hasNext && e == nil {
		e = EmptyIteratorError
	}
	it.hasNext = it.scanner.Scan()
	return
}

const jabberwocky string = `
He took his vorpal sword in hand;
  Long time the manxome foe he sought—
So rested he by the Tumtum tree,
  And stood awhile in thought.
`
