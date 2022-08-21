package main

import (
	"fmt"
	"github.com/samber/lo"
)

func add(x int, y int) int {
	return x + y
}

// Named Return Values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func printV() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// fibonacci is a function that returns
// a function that returns an int.
// Implement a fibonacci function that returns a function (a closure) that
// returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
func fibonacci() func() int {
	count := 0
	prev := 1
	prev2 := 0
	return func() int {
		var result int = 0
		switch count {
		case 0:
			result = prev2
		case 1:
			result = prev
		default:
			result = prev + prev2
			prev2, prev = prev, result
		}
		count = count + 1
		return result
	}
}

func main() {
	fmt.Println("hello world")
	names := lo.Uniq[string]([]string{"Samuel", "Marc", "Samuel"})
	fmt.Println(names)
	//fmt.Println(add(42, 13))
	fmt.Println(split(17))
	printV()
	/*
		f := fibonacci()
		for i := 0; i < 20; i++ {
					fmt.Println(f())
		}
	*/
	TestCon()
}
