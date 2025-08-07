# gogl - Go Generic Library

Let's be frank, we have been anoyed by the lack of standard library features inside the go language. Well this repo is a personal library to try to fill some of those voids.

With the addition of generics and iterators, some of those features can finaly be implemented without having to ressorts to `interface{}` and an expensive use of reflexion to do any sort of generic work. Implementing those missing standard features is now less of an hassle and can be done in a more elegant manner.

## Set

A `Set` type to be more expressive about the developper's intent then having a `map[T]struct{}`

### New

Creates a new empty `Set`.

### NewWithValues

Creates a new `Set` with the provided values.

### Len

Returns the number of elements in the `Set`.

### IsEmpty

Checks if the `Set` is empty.

### Insert

Inserts a value into the `Set`. If the value already exists, it does nothing.

### Contains

Checks if a value is in the `Set`.

### Remove

Removes a value from the `Set`.

### Clear

Clears all values from the `Set`.

### All

Returns an iterator over all elements in the `Set`.

## Slices

A suite of helper function for slices

### Compare

Compares 2 slices and checks if all the values are equals

### CompareN

Compares 2 slices and checks if the first N values are equals

## Iter

A suite of helper iterator functions

### Filter

Filters an iterator based on a predicate function.

### Unzip

Unzips an iterator of pairs into two separate iterators.
