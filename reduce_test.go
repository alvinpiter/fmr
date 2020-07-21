package fmr

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func multiply(accumulator, num int) int {
	return accumulator * num
}

func concat(accumulator, str string) string {
	return fmt.Sprintf("%s %s", accumulator, str)
}

type minMaxSum struct {
	Min int
	Max int
	Sum int
}

func getMinMaxSum(accumulator minMaxSum, num int) minMaxSum {
	if num < accumulator.Min {
		accumulator.Min = num
	}

	if num > accumulator.Max {
		accumulator.Max = num
	}

	accumulator.Sum = accumulator.Sum + num

	return accumulator
}

func TestReduceToInt(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	got := Reduce(slice, multiply, 1)
	want := 120

	assert.Equal(t, want, got)
}

func TestReduceToString(t *testing.T) {
	slice := []string{"I", "am", "so", "hungry"}

	got := Reduce(slice, concat, "")
	want := "I am so hungry"

	assert.Equal(t, want, got)
}

func TestReduceToStruct(t *testing.T) {
	slice := []int{5, 3, 2, 1, 4}

	stats := minMaxSum{Min: 1000, Max: -1000, Sum: 0}

	got := Reduce(slice, getMinMaxSum, stats)
	want := minMaxSum{Min: 1, Max: 5, Sum: 15}

	assert.True(t, reflect.DeepEqual(got, want), fmt.Sprintf("TestReduceToStruct fails, got: %v, want: %v", got, want))
}

func TestReduceWithInvalidSlice(t *testing.T) {
	slice := "this is actually a string"

	assert.Panics(t, func() { Reduce(slice, multiply, 0) }, "TestReduceWithInvalidSlice did not panic")
}

func TestReduceWithInvalidFunc(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	assert.Panics(t, func() { Reduce(slice, concat, "") }, "TestReduceWithInvalidFunc did not panic")
	assert.Panics(t, func() { Reduce(slice, multiply, "") }, "TestReduceWithInvalidFunc did not panic")
	assert.Panics(t, func() { Reduce(slice, "not a func", 0) }, "TestReduceWithInvalidFunc did not panic")
}
