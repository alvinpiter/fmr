package fmr

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type person struct {
	name string
	age  int
}

func isEvenInt(num int) bool {
	if num%2 == 0 {
		return true
	}

	return false
}

func isEvenLengthString(str string) bool {
	if len(str)%2 == 0 {
		return true
	}

	return false
}

func isPersonOver50(p person) bool {
	if p.age > 50 {
		return true
	}

	return false
}

func invalidFunc(num int) {
	num = num * 2
}

func TestFilterWithSliceOfIntegers(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//Filter even numbers
	got := Filter(slice, isEvenInt)
	want := []int{2, 4, 6, 8, 10}

	assert.True(t, reflect.DeepEqual(got, want), fmt.Sprintf("TestFilterWithSliceOfIntegers fails, got: %v, want: %v", got, want))
}

func TestFilterWithSliceOfStrings(t *testing.T) {
	slice := []string{"even", "odd", "eveneven", "oddoddodd"}

	//Filter even-length strings
	got := Filter(slice, isEvenLengthString)
	want := []string{"even", "eveneven"}

	assert.True(t, reflect.DeepEqual(got, want), fmt.Sprintf("TestFilterWithSliceOfStrings fails, got: %v, want: %v", got, want))
}

func TestFilterWithSliceOfStructs(t *testing.T) {
	slice := []person{
		person{name: "Alice", age: 50},
		person{name: "Bob", age: 51},
		person{name: "Charlie", age: 26},
	}

	//Filter person whose age > 50
	got := Filter(slice, isPersonOver50)
	want := []person{person{name: "Bob", age: 51}}

	assert.True(t, reflect.DeepEqual(got, want), fmt.Sprintf("TestFilterWithSliceOfStructs fails, got: %v, want: %v", got, want))
}

func TestFilterWithInvalidSlice(t *testing.T) {
	slice := "this is actually a string"

	assert.Panics(t, func() { Filter(slice, isEvenInt) }, "TestFilterWithInvalidSlice did not panic")
}

func TestFilterWithInvalidFunc(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	assert.Panics(t, func() { Filter(slice, invalidFunc) }, "TestFilterWithInvalidFunc did not panic")
}

func TestFilterWithMismatchType(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	assert.Panics(t, func() { Filter(slice, isEvenLengthString) }, "TestFilterWithMismatchType did not panic")
}
