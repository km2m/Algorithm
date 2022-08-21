package main

import "fmt"

type step struct {
	x int
	y int
}

func TestCon() {
	//steps := []int{1, 2, 3, 4}
	steps := []step{
		{1, 0},
		{0, -1},
		{-1, 0},
		{0, 1}}
	for _, a := range steps {
		fmt.Println(a.x, a.y)
	}
	fmt.Println(steps[3])
	fmt.Println("TestCon Called")
	fmt.Println(steps)
}
