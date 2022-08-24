package leetcode

import "fmt"

func twoSum(num []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(num); i++ {
		another := target - num[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[num[i]] = i
	}
	return nil
}

type ListNode struct {
	val  int
	Next *ListNode
}

func (n *ListNode) Show() (rs string) {
	rs = ""
	for i := n; i != nil; i = i.Next {
		rs = rs + fmt.Sprintf("%d->", i.val)
	}
	return
}

func (n *ListNode) TotalVal() (rs int) {
	base := 1
	rs = 0
	for ; n != nil; n = n.Next {
		rs += n.val * base
		base = base * 10
	}
	return
}

func MakeListNode(num int) *ListNode {
	s := new(ListNode)
	i := s
	for (num / 10) != 0 {
		i.val = num % 10
		num = num / 10
		i.Next = new(ListNode)
		i = i.Next
	}
	i.val = num
	i.Next = nil
	return s
}

// here we abstract The ListNode As Integer
func addTwoNumbers0(l1 *ListNode, l2 *ListNode) *ListNode {
	return MakeListNode(l1.TotalVal() + l2.TotalVal())
	//return nil
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}
