package leetcode

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TwoSum function", func() {
	It("Should Resolve TwoSum", func() {
		//snailMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		Expect(twoSum([]int{1, 2, 3, 4, 5, 8, 9}, 12)).To(Equal([]int{3, 5}))
	})
})

func TestAddTwoNumbers(t *testing.T) {
	n1 := MakeListNode(99991)
	n2 := MakeListNode(19998)
	t.Log(addTwoNumbers(n1, n2).Show())
	t.Log(addTwoNumbers0(n1, n2).Show())
	//t.Log(n.TotalVal())
	t.Log(MakeListNode(999991).Show())
	//t.Errorf("hello world")
}
