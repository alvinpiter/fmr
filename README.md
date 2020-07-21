[![Documentation](https://godoc.org/github.com/alvinpiter/fmr?status.svg)](http://godoc.org/github.com/alvinpiter/fmr)
[![Go Report Card](https://goreportcard.com/badge/github.com/alvinpiter/fmr)](https://goreportcard.com/report/github.com/alvinpiter/fmr)

# fmr

`fmr` is a package that provides Javascript's filter, map, and reduce implemented in Go.

# Examples

## Filter

```
func main() {
	//Filter even numbers
	sliceInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(fmr.Filter(sliceInt, func(num int) bool {
		if num%2 == 0 {
			return true
		}

		return false
	}))

	//Filter string whose length is even
	sliceStr := []string{"even", "odd", "eveneven", "oddoddodd"}
	fmt.Println(fmr.Filter(sliceStr, func(str string) bool {
		if len(str)%2 == 0 {
			return true
		}

		return false
	}))

	type person struct {
		name string
		age  int
	}

	//Filter people whose age is > 50
	sliceStruct := []person{
		person{name: "Alice", age: 50},
		person{name: "Bob", age: 51},
		person{name: "Charlie", age: 26},
	}
	fmt.Println(fmr.Filter(sliceStruct, func(p person) bool {
		if p.age > 50 {
			return true
		}

		return false
	}))
}
```

The above code will output:
```
[2 4 6 8 10]
[even eveneven]
[{Bob 51}]
```

## Map

```
func main() {
	sliceInt := []int{1, 2, 3, 4, 5}

	//Multiply each element by two
	fmt.Println(fmr.Map(sliceInt, func(num int) int {
		return 2 * num
	}))

	//Get parity of each element
	fmt.Println(fmr.Map(sliceInt, func(num int) string {
		if num%2 == 0 {
			return "even"
		}

		return "odd"
	}))

	type person struct {
		name string
		age  int
	}

	names := []string{"Alvin", "Teddy"}

	//Convert each string into person object
	fmt.Println(fmr.Map(names, func(name string) person {
		return person{
			name: name,
			age:  0,
		}
	}))
}
```

The above code will output:
```
[2 4 6 8 10]
[odd even odd even odd]
[{Alvin 0} {Teddy 0}]
```

## Reduce

```
func main() {
	sliceInt := []int{1, 2, 3, 4, 5}

	//Multiply all the numbers
	fmt.Println(fmr.Reduce(sliceInt, func(accumulator, num int) int {
		return accumulator * num
	}, 1))

	sliceString := []string{"alvin", "piter"}

	//Concat all the strings
	fmt.Println(fmr.Reduce(sliceString, func(accumulator, str string) string {
		return fmt.Sprintf("%s %s", accumulator, str)
	}, ""))

	type minMaxSum struct {
		Min int
		Max int
		Sum int
	}

	numbers := []int{10, 1, 5, 4, 3, 2, 7, 6, 9, 8}
	mms := minMaxSum{1000, -1000, 0}

	//Find min, max and sum of numbers
	fmt.Println(fmr.Reduce(numbers, func(accumulator minMaxSum, num int) minMaxSum {
		if num < accumulator.Min {
			accumulator.Min = num
		}

		if num > accumulator.Max {
			accumulator.Max = num
		}

		accumulator.Sum += num

		return accumulator
	}, mms))
}
```

The above code will output:
```
120
 alvin piter
{1 10 55}
```
