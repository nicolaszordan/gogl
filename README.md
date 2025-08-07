# gogl - Go Generic Library

Let's be frank, we have been anoyed by the lack of standard library features inside the go language. Well this repo is a personal library to try to fill some of those voids.

With the addition of generics and iterators, some of those features can finaly be implemented without having to ressorts to `interface{}` and an expensive use of reflexion to do any sort of generic work. Implementing those missing standard features is now less of an hassle and can be done in a more elegant manner.

## Set

A `set` type to be more expressive about the developper's intent then having a `map[T]struct{}`

## Slices

A suite of helper function for slices

## Iter

A suite of helper iterator functions
