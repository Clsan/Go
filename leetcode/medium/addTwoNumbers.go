package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{Val: -1}
	curr := res

	magicNumber := 0
	for l1 != nil || l2 != nil {
		tempSum := magicNumber
		if l1 != nil {
			tempSum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tempSum += l2.Val
			l2 = l2.Next
		}
		curr.Val = tempSum % 10
		if tempSum > 9 {
			magicNumber = 1
		} else {
			magicNumber = 0
		}
		if l1 != nil || l2 != nil {
			curr.Next = &ListNode{Val: -1}
			curr = curr.Next
		} else {
			if magicNumber == 1 {
				curr.Next = &ListNode{Val: 1}
				curr = curr.Next
			}
		}
	}
	return res
}

func main() {
	l1 := &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	res := addTwoNumbers(l1, l2)
	for res != nil {
		print(res.Val)
		res = res.Next
	}
}
