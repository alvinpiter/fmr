package fmr

import (
	"fmt"
	"reflect"
)

var boolType = reflect.TypeOf(true)

/*
Filter will return a new slice containing elements from
given slice that fulfills the filterFunc (filterFunc returns true for
that element). For a given slice with type []T, the filterFunc must be
a func that takes type T as an argument an returns a bool. Otherwise,
Filter will panic.
*/
func Filter(slice, filterFunc interface{}) interface{} {
	sliceVal := reflect.ValueOf(slice)
	funcVal := reflect.ValueOf(filterFunc)

	if sliceVal.Kind() != reflect.Slice {
		panic("First argument must be a slice")
	}

	if funcVal.Kind() != reflect.Func {
		panic("Second argument must be a func")
	}

	elemType := reflect.TypeOf(slice).Elem()
	if isValidFilterFunc(funcVal, elemType) == false {
		panic(fmt.Sprintf("filterFunc must receive one parameter with type %s and returns a bool", elemType.Name()))
	}

	filteredIndices := []int{}
	filterParams := [1]reflect.Value{}
	for i := 0; i < sliceVal.Len(); i++ {
		filterParams[0] = sliceVal.Index(i)
		if funcVal.Call(filterParams[:])[0].Bool() == true {
			filteredIndices = append(filteredIndices, i)
		}
	}

	result := reflect.MakeSlice(sliceVal.Type(), len(filteredIndices), len(filteredIndices))
	for i := 0; i < result.Len(); i++ {
		result.Index(i).Set(sliceVal.Index(filteredIndices[i]))
	}

	return result.Interface()
}

func isValidFilterFunc(fn reflect.Value, elemType reflect.Type) bool {
	//fn must receive 1 parameter with type elemType
	if fn.Type().NumIn() != 1 || fn.Type().In(0) != elemType {
		return false
	}

	//fn must returns a bool
	if fn.Type().NumOut() != 1 || fn.Type().Out(0) != boolType {
		return false
	}

	return true
}
