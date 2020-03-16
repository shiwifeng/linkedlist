package mylinked

type CircleLinkedList struct {
	linkedNode
}

// 双向环形构造函数
func NewCircleLinkedList() *CircleLinkedList {
	return &CircleLinkedList{}
}

// 获取index位置对应的节点对象
func (l4 *CircleLinkedList) node(index int) (*nodes, error) {
	err := rangeCheck(index, l4.size)
	if err != nil {
		return nil, err
	}
	// 1. > size 的一半，尾部查找
	if index <= (l4.size >> 1) {
		node := l4.first
		for i := 0; i < index; i++ {
			node = node.next
		}
		return node, nil
	} else {
		node := l4.last
		for i := l4.size - 1; i > index; i-- {
			node = node.prev
		}
		return node, nil
	}
}
