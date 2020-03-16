package mylinked

type singleCircleLinkedList struct {
	singleLinkedNode
}

func NewSingleCircleLinkedList() *singleCircleLinkedList {
	return &singleCircleLinkedList{}
}

// 添加元素
func (l3 *singleCircleLinkedList) Add(index int, args ...interface{}) error {
	// 1. 第一次添加
	if l3.size == 0 {
		l3.first = &node{}
		l3.first.data = args[0]
		l3.size++
		if len(args) > 1 {
			err := l3.append(args[1:], l3.size-1, 1)
			if err != nil {
				l3.Clear()
				return err
			}
		} else {
			l3.first.next = l3.first
		}
		return nil
	}
	// 2. 尾部追加
	if l3.size == index {
		err := l3.append(args, l3.size-1, 1)
		return err
	}
	// 3. 中间插入
	if tmpIndex := l3.size - 1; index > 0 && index <= tmpIndex {
		// 3.1 获取要插入的节点，先保存
		newNode, err := l3.node(index)
		if err != nil {
			return err
		}
		// 3.2 追加当前的父节点后面
		err = l3.append(args, index-1)
		if err != nil {
			return err
		}
		// 3.3 获取刚插入的最后节点，index-1+len(args): 步骤3.1的索引位置-1 + 已经插入的数量
		nextNode, err := l3.node(index + len(args) - 1)
		if err != nil {
			l3.Clear()
			return err
		}
		nextNode.next = newNode
		return nil
	}
	// 4. 第一位插入
	// 4.1 获取要插入的节点，先保存
	newNode, err := l3.node(index)
	if err != nil {
		return err
	}
	// 4.2 更新入口，指向新插入的节点
	l3.first = &node{}
	l3.first.data = args[0]
	l3.size++
	// 4.3 批量追加
	if len(args) > 1 {
		err = l3.append(args[1:], 0)
		if err != nil {
			l3.Clear()
			return err
		}
	}
	// 4.4 获取刚插入的最后节点，index+len(args): 步骤3.1的索引位置 + 已经插入的数量
	nextNode, err := l3.node(index + len(args) - 1)
	if err != nil {
		l3.Clear()
		return err
	}
	nextNode.next = newNode
	// 4.5 修改最后的节点指向插入的第一个
	endNode, err := l3.node(l3.size - 1)
	if err != nil {
		l3.Clear()
		return err
	}
	endNode.next = l3.first
	return nil
}

// 清空链表
func (l3 *singleCircleLinkedList) Clear() {
	l3.first = nil
	l3.size = 0
	return
}

// 根据索引获取元素
func (l3 *singleCircleLinkedList) Get(index int) (interface{}, error) {
	data, err := l3.node(index)
	if err != nil {
		return nil, err
	}
	return data.data, nil
}

// 获取元素的容量
func (l3 *singleCircleLinkedList) Size() int {
	return l3.size
}

// 根据索引设置元素
func (l3 *singleCircleLinkedList) Set(index int, element interface{}) (err error) {
	node, err := l3.node(index)
	if err != nil {
		return
	}
	node.data = element
	return
}

// 获取元素对应的第一次的索引，暂时不支持获取引用类型元素的获取
func (l3 *singleCircleLinkedList) IndexOf(element interface{}) int {
	return singleIndexOf(element, &l3.singleLinkedNode)
}

// 移除元素
func (l3 *singleCircleLinkedList) Remove(index int) (interface{}, error) {
	err := rangeCheck(index, l3.size)
	if err != nil {
		return nil, err
	}
	// 0. 容量为1，直接清空
	if l3.size == 1 {
		element := l3.first.data
		l3.Clear()
		return element, nil
	}
	// 1. 移除第一个元素
	if index == 0 {
		// 0. 保存被移除的元素
		element := l3.first.data
		// 1.1 获取第二个元素，先保存
		// 1.2 入口元素指向 1.1 的元素
		l3.first = l3.first.next
		l3.size--
		// 1.3 最后的元素的next，指向第一个
		endNode, err := l3.node(l3.size - 1)
		if err != nil {
			l3.Clear()
			return nil, err
		}
		endNode.next = l3.first
		return element, nil
	}
	// 2. 移除最后的元素
	if index == l3.size-1 {
		node, err := l3.node(l3.size - 2)
		if err != nil {
			return nil, err
		}
		// 2.0. 保存被移除的元素
		element := node.next.data
		// 2.1 最后的元素的next，指向第一个节点
		node.next = l3.first
		l3.size--
		return element, nil
	}
	// 3. 移除中间的元素
	// 3.1 移除的节点前一个
	firstNode, err := l3.node(index - 1)
	if err != nil {
		return nil, err
	}
	// 3.2 保存移除的元素
	element := firstNode.next.data
	// 3.3 移除的节点后一个
	tailNode := firstNode.next.next
	firstNode.next = tailNode
	l3.size--
	return element, nil
}

func (l3 *singleCircleLinkedList) append(args []interface{}, index int, a ...interface{}) error {
	nextNode, err := l3.node(index) // 获取节点
	if err != nil {
		return err
	}
	isEndNode := len(a) < 1
	for i, v := range args {
		newNode := &node{}
		// 1. 最后一个指向第一个，形成闭环
		if i == len(args)-1 && isEndNode {
			newNode.next = l3.first
		}
		newNode.data = v
		nextNode.next = newNode // 把当前的节点赋值父节点
		nextNode = newNode      // 把当前的保存，让下一次循环变成父节点
		l3.size++
	}
	return nil
}

// 获取index位置对应的节点对象
func (l3 *singleCircleLinkedList) node(index int) (*node, error) {
	err := rangeCheck(index, l3.size)
	if err != nil {
		return nil, err
	}
	node := l3.first
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node, nil
}
