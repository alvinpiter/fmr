package fmr

import (
	"fmt"
	"reflect"
)

/*
Reduce will return a value which is the result of reducing all elements
in the given slice by using reduceFunc. For a given slice with type []T and initialValue
with type T`, the reduceFunc must accept two parameters. The first one is `accumulator` (with type T') and
the second one is a parameter with type T. The type returned by reduceFunc must be T'.
Otherwise, Reduce will panic.
*/
func Reduce(slice, reduceFunc, initialValue interface{}) interface{} {
	sliceVal := reflect.ValueOf(slice)
	funcVal := reflect.ValueOf(reduceFunc)
	initialVal := reflect.ValueOf(initialValue)

	if sliceVal.Kind() != reflect.Slice {
		panic("First argument must be a slice")
	}

	if funcVal.Kind() != reflect.Func {
		panic("Second argument must be a func")
	}

	elemType := reflect.TypeOf(slice).Elem()
	reduceType := reflect.TypeOf(initialValue)
	if isValidReduceFunc(funcVal, elemType, reduceType) == false {
		panic(fmt.Sprintf("reduceFunc must receive 2 parameters each of type %s and %s. And return a type %s", reduceType, elemType, reduceType))
	}

	if sliceVal.Len() == 0 {
		return initialValue
	}

	reduceParams := [2]reflect.Value{}
	reduceValue := initialVal
	for i := 0; i < sliceVal.Len(); i++ {
		reduceParams[0] = reduceValue
		reduceParams[1] = sliceVal.Index(i)
		reduceValue = funcVal.Call(reduceParams[:])[0]
	}

	return reduceValue.Interface()
}

func isValidReduceFunc(fn reflect.Value, elemType, reduceType reflect.Type) bool {
	//fn must receive 2 parameters with type reduceType and elemType
	if fn.Type().NumIn() != 2 || fn.Type().In(0) != reduceType || fn.Type().In(1) != elemType {
		return false
	}

	//fn must return type reduceType
	if fn.Type().NumOut() != 1 || fn.Type().Out(0) != reduceType {
		return false
	}

	return true
}
