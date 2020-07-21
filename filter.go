package fmr

/*
Filter will return a new slice containing elements from
given slice that fulfills the filterFunc (filterFunc returns true for
that element). For a given slice with type []T, the filterFunc must be
a func that takes type T as an argument an returns a bool. Otherwise,
Filter will panic.
*/
func Filter(slice, filterFunc interface{}) interface{} {
	return nil
}
