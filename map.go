package fmr

/*
Map will return a new slice with length equal to length of given slice.
Each element of the new slice is the result of applying mapFunc to
the element at the same index in given slice. For a given slice of type []T,
mapFunc must be a func that takes type T as an argument and returns exactly 1 value.
Otherwise, Map will panic.
*/
func Map(slice, mapFunc interface{}) interface{} {
	return nil
}
