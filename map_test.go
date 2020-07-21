package fmr

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func multiplyIntByTwo(num int) int {
	return 2 * num
}

func parity(num int) string {
	if num%2 == 0 {
		return "even"
	}

	return "odd"
}

func sliceToPerson(slice []interface{}) person {
	return person{
		name: slice[0].(string),
		age:  slice[1].(int),
	}
}

func TestMapToSliceOfIntegers(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	//Multiply each number by 2
	got := Map(slice, multiplyIntByTwo)
	want := []int{2, 4, 6, 8, 10}

	assert.True(t, reflect.DeepEqual(got, want), fmt.Sprintf("TestMapToSliceOfIntegers fails, got: %v, want: %v", got, want))
}

func TestMapToSliceOfStrings(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	//Map each number to its parity
	got := Map(slice, parity)
	want := []string{"odd", "even", "odd", "even", "odd"}

	assert.True(t, reflect.DeepEqual(got, want), fmt.Sprintf("TestmapToSliceOfStrings fails, got: %v, want: %v", got, want))
}

func TestMapToSliceOfStructs(t *testing.T) {
	slice := []interface{}{
		[]interface{}{"Alvin", 26},
		[]interface{}{"Teddy", 23},
	}

	got := Map(slice, sliceToPerson)
	want := []person{
		person{name: "Alvin", age: 26},
		person{name: "Teddy", age: 23},
	}

	assert.True(t, reflect.DeepEqual(got, want), fmt.Sprintf("TestMapToSliceOfStructs fails, got: %v, want: %v", got, want))
}

func TestMapWithInvalidSlice(t *testing.T) {
	slice := "this is actually a string"

	assert.Panics(t, func() { Map(slice, parity) }, "TestMapWithInvalidSlice did not panic")
}

func TestMapWithInvalidFunc(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	assert.Panics(t, func() { Map(slice, invalidFunc) }, "TestMapWithInvalidFunc did not panic")
	assert.Panics(t, func() { Map(slice, "not a func") }, "TestMapWithInvalidFunc did not panic")
}

func TestMapWithMismatchType(t *testing.T) {
	slice := []string{"string1", "string2"}

	assert.Panics(t, func() { Map(slice, parity) }, "TestMapWithMismatchType did not panic")
}
