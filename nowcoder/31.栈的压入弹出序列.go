package nowcoder

import (
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。
//假设压入栈的所有数字均不相等。
//例如序列 1,2,3,4,5 是某栈的压入顺序，序列 4,5,3,2,1 是该压栈序列对应的一个弹出序列，
//但 4,3,5,1,2 就不可能是该压栈序列的弹出序列。(注：两个序列长度相等)

func isPopOrder(pushOrder, popOrder []int) bool {
	s := new(datastructures.Stack)
	for pushIndex, popIndex, n := 0, 0, len(pushOrder); pushIndex < n; pushIndex++ {
		s.Push(pushOrder[pushIndex])
		for popIndex < n && !s.Empty() && popOrder[popIndex] == s.Peek() {
			s.Pop()
			popIndex++
		}
	}
	return s.Empty()
}
