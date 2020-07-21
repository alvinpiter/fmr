package fmr

/*
Reduce will return a value which is the result of reducing all elements
in the given slice by using reduceFunc. For a given slice with type []T and initialValue
with type T`, the reduceFunc must accept two parameters. The first one is`accumulator` (with type T') and
the second one is a parameter with type T. The return value of the reduceFunc must be T'.
Otherwise, Reduce will panic.
*/
func Reduce(slice, reduceFunc, initialValue interface{}) interface{} {
	return nil
}
