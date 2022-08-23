package leetcode

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TwoSum function", func() {
	It("Should Resolve TwoSum", func() {
		//snailMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		Expect(twoSum([]int{1, 2, 3, 4, 5, 8, 9}, 12)).To(Equal([]int{3, 5}))
	})
})
