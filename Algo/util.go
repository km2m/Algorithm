package Algo

import (
	"bytes"
	"errors"
	"math/rand"
	"time"
)

//import "golang.org/x/exp/rand"

// r := rand.Rand{}
// println(RandString(&r, 5, "abc"))
func randString(size int, alphabet string) string {
	var buffer bytes.Buffer
	for i := 0; i < size; i++ {
		index := rand.Intn(len(alphabet))
		buffer.WriteString(string(alphabet[index]))
	}
	return buffer.String()
}

func RandString() string {
	//r := rand.Rand{}
	//println(randString(&r, 5, "abc"))
	rand.Seed(time.Now().UnixNano())
	alpahabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rs := randString(rand.Intn(10)+6, alpahabet)
	//fmt.Println(rs)
	return rs
	//return "123"
}

func MakeRange(begin int, end int, step int) (sequence []int) {
	if step == 0 {
		panic(errors.New("step must not be zero"))
	}
	count := 0
	if (end > begin && step > 0) || (end < begin && step < 0) {
		count = (end-step-begin)/step + 1
	}

	sequence = make([]int, count)
	for i := 0; i < count; i, begin = i+1, begin+step {
		sequence[i] = begin
	}
	return
}
