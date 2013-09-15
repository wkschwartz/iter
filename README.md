iter
====

Go package `iter` demonstrates an iterator interface and protocol.

Because Go doesn't have generics, you'll have to write your own version of the
interfaces for your data types. Thus this package is mainly useful for the
documentation and examples, which you can run with `go test`. Think of `iter` as
a protocol definition.

Iterators should not replace the `for x := range` statement, which is vastly
easier to use. However I see them as a way to deal with more complex data
types. In the long run, I think it would be nice if the `range` statement took
an iterator like Java's foreach or Python's `for x in` do. In the mean time, I'm
going to try to use them in my code to make iteration more standardized and my
APIs easier to remember and read.

Please send feedback by posting a bug report or patches through pull
requests. I'd love to hear what you think!

I'm releasing this code under the BSD license in the LICENSE file so you can use
`iter` in your projects.