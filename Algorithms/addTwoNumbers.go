package algorithms

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resListNode := &ListNode{}
	cur := resListNode
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return resListNode.Next
}

// 执行时间最少写法
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	ca := 0
// 	return addTwoNumber(l1, l2, &ca)
// }

// func addTwoNumber(l1 *ListNode, l2 *ListNode, carry *int) (l3 *ListNode) {
// 	// 退出条件：两个节点都是nil，且进位为0
// 	if l1 == nil && l2 == nil && *carry == 0 {
// 		return nil
// 	}
// 	// 算进位
// 	sum := *carry
// 	if l1 != nil {
// 		sum += l1.Val
// 		l1 = l1.Next
// 	} else {
// 		l1 = nil
// 	}
// 	if l2 != nil {
// 		sum += l2.Val
// 		l2 = l2.Next
// 	} else {
// 		l2 = nil
// 	}

// 	rem := sum % 10
// 	*carry = sum / 10

// 	return &ListNode{
// 		Val:  rem,
// 		Next: addTwoNumber(l1, l2, carry),
// 	}
// }

// 消耗内存最小的写法
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	next := 0
// 	h := l1
// 	for l1 != nil && l2 != nil {
// 		op1 := l1.Val
// 		op2 := l2.Val
// 		l1.Val, next = (op1+op2+next)%10, (op1+op2+next)/10
// 		if l1.Next != nil && l2.Next != nil {
// 			l1 = l1.Next
// 			l2 = l2.Next
// 			continue
// 		}
// 		if l1.Next == nil {
// 			l1.Next = l2.Next
// 		}
// 		for l1.Next != nil && next > 0 {
// 			l1 = l1.Next
// 			l1.Val, next = (l1.Val+next)%10, (l1.Val+next)/10
// 		}
// 		break
// 	}
// 	if next > 0 {
// 		l1.Next = &ListNode{next, nil}
// 	}
// 	return h
// }
