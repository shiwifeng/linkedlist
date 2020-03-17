package linkedlist

type SingleLinkedList struct {
	singleLinkedNode
}

// 单向链表构造函数
func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{}
}

// 添加元素
func (l1 *SingleLinkedList) Add(index int, args ...interface{}) error {
	// 1. 第一次添加
	if l1.size == 0 {
		l1.first = &node{}
		l1.first.data = args[0]
		l1.size++
		if len(args) > 1 {
			err := l1.append(args[1:], l1.size-1)
			if err != nil {
				l1.Clear()
				return err
			}
		}
		return nil
	}
	// 2. 尾部追加
	if l1.size == index {
		err := l1.append(args, l1.size-1)
		return err
	}
	// 3. 中间插入
	if tmpIndex := l1.size - 1; index > 0 && index <= tmpIndex {
		// 3.1 获取要插入的节点，先保存
		newNode, err := l1.node(index)
		if err != nil {
			return err
		}
		// 3.2 追加当前的父节点后面
		err = l1.append(args, index-1)
		if err != nil {
			return err
		}
		// 3.3 获取刚插入的最后节点，index-1+len(args): 步骤3.1的索引位置-1 + 已经插入的数量
		nextNode, err := l1.node(index + len(args) - 1)
		if err != nil {
			l1.Clear()
			return err
		}
		nextNode.next = newNode
		return nil
	}
	// 4. 第一位插入
	// 4.1 获取要插入的节点，先保存
	newNode, err := l1.node(index)
	if err != nil {
		return err
	}
	// 4.2 更新入口，指向新插入的节点
	l1.first = &node{}
	l1.first.data = args[0]
	l1.size++
	// 4.3 批量追加
	if len(args) > 1 {
		err = l1.append(args[1:], 0)
		if err != nil {
			l1.Clear()
			return err
		}
	}
	// 4.4 获取刚插入的最后节点，index+len(args): 步骤3.1的索引位置 + 已经插入的数量
	nextNode, err := l1.node(index + len(args) - 1)
	if err != nil {
		l1.Clear()
		return err
	}
	nextNode.next = newNode
	return nil
}

func (l1 *SingleLinkedList) append(args []interface{}, index int) error {
	nextNode, err := l1.node(index) // 获取节点
	if err != nil {
		return err
	}
	for _, v := range args {
		newNode := &node{}
		newNode.data = v
		nextNode.next = newNode // 把当前的节点赋值父节点
		nextNode = newNode      // 把当前的保存，让下一次循环变成父节点
		l1.size++
	}
	return nil
}

// 清除链表
func (l1 *SingleLinkedList) Clear() {
	l1.first = nil
	l1.size = 0
	return
}

// 获取元素
func (l1 *SingleLinkedList) Get(index int) (interface{}, error) {
	data, err := l1.node(index)
	if err != nil {
		return nil, err
	}
	return data.data, nil
}

// 获取元素对应的第一次的索引，暂时不支持获取引用类型元素的获取
func (l1 *SingleLinkedList) IndexOf(element interface{}) int {
	return singleIndexOf(element, &l1.singleLinkedNode)
}

// 获取元素的容量
func (l1 *SingleLinkedList) Size() int {
	return l1.size
}

// 根据索引设置元素
func (l1 *SingleLinkedList) Set(index int, element interface{}) (interface{}, error) {
	node, err := l1.node(index)
	if err != nil {
		return nil, err
	}
	node.data = element
	return element, nil
}

// 移除元素
func (l1 *SingleLinkedList) Remove(index int) (interface{}, error) {
	err := rangeCheck(index, l1.size)
	if err != nil {
		return nil, err
	}
	// 0. 容量为1，直接清空
	if l1.size == 1 {
		element := l1.first.data
		l1.Clear()
		return element, nil
	}
	// 1. 移除第一个元素
	if index == 0 {
		element := l1.first.data
		// 1.1 获取第二个元素，先保存
		// 1.2 入口元素指向 1.1 的元素
		l1.first = l1.first.next
		l1.size--
		return element, nil
	}
	// 2. 移除最后的元素
	if index == l1.size-1 {
		node, err := l1.node(l1.size - 2)
		if err != nil {
			return nil, err
		}
		element := node.next.data
		node.next = nil
		l1.size--
		return element, nil
	}
	// 3. 移除中间的节点
	// 3.1 移除的节点前一个
	firstNode, err := l1.node(index - 1)
	if err != nil {
		return nil, err
	}
	// 3.2 移除的节点后一个
	tailNode := firstNode.next.next
	// 3.3 保存移除的元素
	element := firstNode.data
	firstNode.next = tailNode
	l1.size--
	return element, nil
}

// 获取index位置对应的节点对象
func (l1 *SingleLinkedList) node(index int) (*node, error) {
	err := rangeCheck(index, l1.size)
	if err != nil {
		return nil, err
	}
	node := l1.first
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node, nil
}
