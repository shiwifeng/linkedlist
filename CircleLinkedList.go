package linkedlist

type CircleLinkedList struct {
	linkedNode
}

// 双向环形构造函数
func NewCircleLinkedList() *CircleLinkedList {
	return &CircleLinkedList{}
}

// 添加元素
func (l4 *CircleLinkedList) Add(index int, args ...interface{}) error {
	// 1. 第一次添加
	if l4.size == 0 {
		l4.first = &nodes{}
		l4.first.data = args[0]
		l4.first.next = l4.first
		l4.first.prev = l4.first
		l4.last = &nodes{}
		l4.last.data = l4.first.data
		l4.last.next = l4.last
		l4.last.prev = l4.last
		l4.size++
		if len(args) > 1 {
			// 1. 获取上一个节点
			frontNode, err := l4.node(l4.size - 1) // 获取节点
			if err != nil {
				return err
			}
			args := args[1:]
			for i, v := range args {
				// 1.1 构建节点
				newNode := &nodes{}
				newNode.data = v
				// 1.2 构建节点指向步骤1.的节点
				newNode.prev = frontNode
				// 1.2.1 最后一个节点的 enxt 指向入口节点第一个,
				// 入口节点的 prev 指向最后一个节点，形成双闭环
				// 出口节点 指向最后一个节点
				if i == len(args)-1 {
					newNode.next = l4.first
					l4.first.prev = newNode
					l4.last = newNode
					//endNode = newNode
				}
				// 1.3 步骤1.的节点的下一个节点指向刚构建的节点
				frontNode.next = newNode
				// 1.4 把当前的保存，让下一次循环变成父节点
				frontNode = newNode
				l4.size++
			}
		}
		return nil
	}
	// 2. 尾部追加
	if l4.size == index {
		// 1. 获取上一个节点
		frontNode, err := l4.node(l4.size - 1) // 获取节点
		if err != nil {
			return err
		}
		for i, v := range args {
			// 1.1 构建节点
			newNode := &nodes{}
			newNode.data = v
			// 1.2 构建节点指向步骤1.的节点
			newNode.prev = frontNode
			// 1.2.1 最后一个节点的 enxt 指向入口节点第一个,
			// 入口节点的 prev 指向最后一个节点，形成双闭环
			// 出口节点 指向最后一个节点
			if i == len(args)-1 {
				newNode.next = l4.first
				l4.first.prev = newNode
				l4.last = newNode
				//endNode = newNode
			}
			// 1.3 步骤1.的节点的下一个节点指向刚构建的节点
			frontNode.next = newNode
			// 1.4 把当前的保存，让下一次循环变成父节点
			frontNode = newNode
			l4.size++
		}
		return err
	}
	// 3. 中间插入
	if tmpIndex := l4.size - 1; index > 0 && index <= tmpIndex {
		// 3.1 获取要插入的节点，先保存
		nextNode, err := l4.node(index)
		if err != nil {
			return err
		}
		// 3.1.1 获取插入位置的上一个节点
		frontNode := nextNode.prev
		for i, v := range args {
			// 1.1 构建节点
			newNode := &nodes{}
			newNode.data = v
			// 1.2 构建节点指向步骤1.的节点
			newNode.prev = frontNode
			if i == len(args)-1 {
				newNode.next = nextNode
				nextNode.prev = newNode
			}
			// 1.3 步骤1.的节点的下一个节点指向刚构建的节点
			frontNode.next = newNode
			// 1.4 把当前的保存，让下一次循环变成父节点
			frontNode = newNode
			l4.size++

		}
		return nil
	}
	// 4. 第一位插入
	// 4.1 获取要首节点，先保存
	nextNode := l4.first
	frontNode := &nodes{}
	for i, v := range args {
		// 1.1 构建节点
		newNode := &nodes{}
		newNode.data = v
		newNode.prev = frontNode
		// 1.2 构建的是第一个，prev 指向最后一个节点，
		// 最后一个节点的 enxt 指向入口节点第一个,形成双闭环
		if i == 0 {
			newNode.prev = l4.last
			l4.last.next = newNode
			l4.first = newNode
		} else {
			if i == len(args)-1 {
				newNode.next = nextNode
				nextNode.prev = newNode
			}
		}
		// 1.3 步骤1.的节点的下一个节点指向刚构建的节点
		frontNode.next = newNode
		// 1.4 把当前的保存，让下一次循环变成父节点
		frontNode = newNode
		l4.size++
	}
	return nil
}

// 清除链表
func (l4 *CircleLinkedList) Clear() {
	l4.first = nil
	l4.size = 0
	return
}

// 获取元素对应的第一次的索引，暂时不支持获取引用类型元素的获取
func (l4 *CircleLinkedList) IndexOf(element interface{}) int {
	return indexOf(element, &l4.linkedNode)
}

// 获取元素
func (l4 *CircleLinkedList) Get(index int) (interface{}, error) {
	data, err := l4.node(index)
	if err != nil {
		return nil, err
	}
	return data.data, nil
}

// 获取元素的容量
func (l4 *CircleLinkedList) Size() int {
	return l4.size
}

// 根据索引设置元素
func (l4 *CircleLinkedList) Set(index int, element interface{}) (interface{}, error) {
	node, err := l4.node(index)
	if err != nil {
		return nil, err
	}
	node.data = element
	return element, err
}

// 移除元素
func (l4 *CircleLinkedList) Remove(index int) (interface{}, error) {
	err := rangeCheck(index, l4.size)
	if err != nil {
		return nil, err
	}
	// 0. 容量为1，直接清空
	if l4.size == 1 {
		element := l4.first.data
		l4.Clear()
		return element, nil
	}
	// 1. 移除第一个节点
	if index == 0 {
		// 0. 保存被移除的元素
		element := l4.first.data
		// 1.1 获取第二个节点，先保存 node
		node := l4.first.next
		// 2. 最尾元素的 enxt 指向 node
		// 入口元素指向 node
		// node 的 prev 指向 last 节点
		node.prev = l4.last
		l4.last.next = node
		l4.first = node
		l4.size--
		return element, nil
	}
	// 2. 移除最后的节点
	if index == l4.size-1 {
		// 0. 保存被移除的元素
		element := l4.last.data
		// 1. 获取移除的上一个节点 node
		node := l4.last.prev
		// 2. node 的 next 指向 头节点
		node.next = l4.first
		l4.first.prev = node
		l4.last = node
		l4.size--
		return element, nil
	}
	// 3. 移除中间的元素
	// 0. 保存被移除的元素
	node, err := l4.node(index)
	if err != nil {
		return nil, err
	}
	firstNode := node.prev
	tailNode := node.next
	firstNode.next = tailNode
	tailNode.prev = firstNode
	l4.size--
	return node.data, nil
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
