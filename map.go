package fmr

import (
	"fmt"
	"reflect"
)

/*
Map will return a new slice with length equal to length of given slice.
Each element of the new slice is the result of applying mapFunc to
the element at the same index in given slice. For a given slice of type []T,
mapFunc must be a func that takes type T as an argument and returns exactly 1 value.
Otherwise, Map will panic.
*/
func Map(slice, mapFunc interface{}) interface{} {
	sliceVal := reflect.ValueOf(slice)
	funcVal := reflect.ValueOf(mapFunc)

	if sliceVal.Kind() != reflect.Slice {
		panic("First argument must be a slice")
	}

	if funcVal.Kind() != reflect.Func {
		panic("Second argument must be a func")
	}

	elemType := reflect.TypeOf(slice).Elem()
	if isValidMapFunc(funcVal, elemType) == false {
		panic(fmt.Sprintf("mapFunc must receive one parameter with type %s and return a value", elemType.Name()))
	}

	funcReturnType := funcVal.Type().Out(0)

	result := reflect.MakeSlice(reflect.SliceOf(funcReturnType), sliceVal.Len(), sliceVal.Len())
	mapParams := [1]reflect.Value{}
	for i := 0; i < result.Len(); i++ {
		mapParams[0] = sliceVal.Index(i)
		result.Index(i).Set(funcVal.Call(mapParams[:])[0])
	}

	return result.Interface()
}

func isValidMapFunc(fn reflect.Value, elemType reflect.Type) bool {
	//fn must receive one parameter with type elemType
	if fn.Type().NumIn() != 1 || fn.Type().In(0) != elemType {
		return false
	}

	//fn must return a value
	if fn.Type().NumOut() != 1 {
		return false
	}

	return true
}
