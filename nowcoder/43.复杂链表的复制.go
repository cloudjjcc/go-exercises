package nowcoder

//题目描述
//输入一个复杂链表（每个节点中有节点值，以及两个指针，一个指向下一个节点，
//另一个特殊指针指向任意一个节点），返回结果为复制后复杂链表的head。
//（注意，输出结果中请不要返回参数中的节点引用，否则判题程序会直接返回空

type randomListNode struct {
	next, random *randomListNode
	value        int
}

func clone(root *randomListNode) *randomListNode {
	// clone node
	tmp := make(map[*randomListNode]*randomListNode)
	cur := root
	for cur != nil {
		tmp[cur] = &randomListNode{value: root.value}
		cur = cur.next
	}
	// clone pointer
	cur = root
	for cur != nil {
		tmp[cur].next = tmp[cur.next]
		tmp[cur].random = tmp[cur.random]
		cur = cur.next
	}
	return tmp[root]
}
