package Algo

import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sort function", func() {
	It("Should Resolve TwoSum", func() {
		//snailMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		//Expect(TwoSum([]int{1, 2, 3}, 5)).To(Equal(nil))
		Expect(Sort0(2, 3)).To(Equal(5))
		Expect(Sort0(2, 4)).To(Equal(6))
	})
})

var _ = Describe("Quick Sort", func() {
	It("", func() {
		Expect(QuickSort([]int{1, 3, 2})).To(Equal([]int{1, 2, 3}))
		Expect(QuickSort(rand.Perm(500))).To(Equal(MakeRange(0, 500, 1)))
		Expect(InsertOrder([]int{1, 3, 2, 0, 4})).To(Equal(MakeRange(0, 5, 1)))
		Expect(InsertOrder(rand.Perm(500))).To(Equal(MakeRange(0, 500, 1)))
		Expect(MergeSort([]int{1, 3, 2, 0, 4})).To(Equal(MakeRange(0, 5, 1)))
		Expect(MergeSort(rand.Perm(500))).To(Equal(MakeRange(0, 500, 1)))
	})
})

func TestQuickSort(t *testing.T) {
	//xs := []int{1, 3, 4, 7, 2, 6}
	slc1 := []int{58, 69, 40, 45, 11, 56, 67, 21, 65}
	slc2 := make([]int, 4)
	copy(slc2, slc1[2:6])
	fmt.Println(slc2)
	//QuickSort(xs, 0, len(xs)-1)
	//sort.Ints(xs)
	//fmt.Println(xs)
	// c := quick.Config{MaxCount: 1000000}
	//if err := quick.Check(addW, &c); err != nil {
	//t.Error(err)
	//}
}
