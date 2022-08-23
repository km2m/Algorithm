package main_test

import (
	"bytes"
	"fmt"
	"golang.org/x/exp/rand"
	"testing"
	"testing/quick"
)

// r := rand.Rand{}
// println(RandString(&r, 5, "abc"))
func RandString(r *rand.Rand, size int, alphabet string) string {
	var buffer bytes.Buffer
	for i := 0; i < size; i++ {
		index := r.Intn(len(alphabet))
		buffer.WriteString(string(alphabet[index]))
	}
	return buffer.String()
}

func add(x, y int) int {
	if x > 20 && x < 100 {
		return x + y + 1
	} else {
		return x + y
	}
}

func addW(x, y int) bool {
	return add(x, y) == (x + y)
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestAdd(t *testing.T) {
	c := quick.Config{MaxCount: 1000000}
	if err := quick.Check(addW, &c); err != nil {
		t.Error(err)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
