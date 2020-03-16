package mylinked

import (
	"reflect"
)

// 双向链表
type linkedList struct {
	linkedNode
}

// 双向链表构造函数
func NewLinkedList() *linkedList {
	return &linkedList{}
}

// 添加元素
func (l2 *linkedList) Add(index int, args ...interface{}) (error) {
	// 1. 第一次添加
	if l2.size == 0 {
		l2.first = &nodes{}
		l2.first.data = args[0]
		l2.last = &nodes{}
		l2.last.data = l2.first.data
		l2.size++
		if len(args) > 1 {
			err := l2.append(args[1:], l2.size-1)
			if err != nil {
				l2.Clear()
				return err
			}
		}
		return nil
	}
	// 2. 尾部追加
	if l2.size == index {
		err := l2.append(args, l2.size-1)
		return err
	}
	// 3. 中间插入
	if tmpIndex := l2.size - 1; index > 0 && index <= tmpIndex {
		// 3.1 获取要插入的节点，先保存
		newNode, err := l2.node(index)
		if err != nil {
			return err
		}
		// 3.2 追加当前的父节点后面
		err = l2.append(args, index-1, 1)
		if err != nil {
			return err
		}
		// 3.3 获取刚插入的最后节点，index-1+len(args): 步骤3.1的索引位置-1 + 已经插入的数量
		nextNode, err := l2.node(index + len(args) - 1)
		if err != nil {
			l2.Clear()
			return err
		}
		newNode.prev = nextNode
		nextNode.next = newNode
		return nil
	}
	// 4. 第一位插入
	// 4.1 获取要插入的节点，先保存
	newNode, err := l2.node(index)
	if err != nil {
		return err
	}
	// 4.2 更新入口，指向新插入的节点
	l2.first = &nodes{}
	l2.first.data = args[0]
	l2.size++
	// 4.3 批量追加
	if len(args) > 1 {
		err = l2.append(args[1:], 0, 1)
		if err != nil {
			l2.Clear()
			return err
		}
	}
	// 4.4 获取刚插入的最后节点，index+len(args): 步骤3.1的索引位置 + 已经插入的数量
	nextNode, err := l2.node(index + len(args) - 1)
	if err != nil {
		l2.Clear()
		return err
	}
	newNode.prev = nextNode
	nextNode.next = newNode
	return nil
}

func (l2 *linkedList) append(args []interface{}, index int, a ...interface{}) error {
	nextNode, err := l2.node(index) // 获取节点
	if err != nil {
		return err
	}
	isLastNode := len(a) < 1
	endNode := &nodes{}
	for i, v := range args {
		newNode := &nodes{}
		newNode.data = v
		newNode.prev = nextNode
		nextNode.next = newNode // 把当前的节点赋值父节点
		nextNode = newNode      // 把当前的保存，让下一次循环变成父节点
		l2.size++
		if i == len(args)-1 && isLastNode {
			endNode = newNode
		}
	}
	if isLastNode {
		l2.last = endNode
	}
	return nil
}

// 清除链表
func (l2 *linkedList) Clear() {
	l2.first = nil
	l2.size = 0
	return
}

// 获取元素
func (l2 *linkedList) Get(index int) (interface{}, error) {
	data, err := l2.node(index)
	if err != nil {
		return nil, err
	}
	return data.data, nil
}

// 获取元素对应的第一次的索引，暂时不支持获取引用类型元素的获取
func (l2 *linkedList) IndexOf(element interface{}) (int) {
	node1 := l2.first
	node2 := l2.last
	lastNum := l2.size
	svg := (l2.size - 1) >> 1 // 边界
	if element == nil {
		for i := 0; i < l2.size; i++ {
			if i > svg {
				break
			}
			lastNum--
			if node1.data == nil {
				return i
			}
			if node2.data == nil {
				return lastNum
			}
			node1 = node1.next
			node2 = node2.prev
		}
	} else {
		for i := 0; i < l2.size; i++ {
			if i > svg {
				break
			}
			lastNum--
			if reflect.DeepEqual(element, node1.data) {
				return i
			}
			if reflect.DeepEqual(element, node2.data) {
				return lastNum
			}
			node1 = node1.next
			node2 = node2.prev
		}
	}
	return ELEMENT_NOT_FOUND
}

// 获取元素的容量
func (l2 *linkedList) Size() int {
	return l2.size
}

// 根据索引设置元素
func (l2 *linkedList) Set(index int, element interface{}) (err error) {
	node, err := l2.node(index)
	if err != nil {
		return
	}
	node.data = element
	return
}

// 移除元素
func (l2 *linkedList) Remove(index int) error {
	err := rangeCheck(index, l2.size)
	if err != nil {
		return err
	}
	// 1. 移除第一个元素
	if index == 0 {
		// 1.1 获取第二个元素，先保存
		// 1.2 入口元素指向 1.1 的元素
		node, err := l2.node(1)
		if err != nil {
			return err
		}
		node.prev = nil
		l2.first = node
		l2.size--
		return nil
	}
	// 2. 移除最后的元素
	if index == l2.size-1 {
		node, err := l2.node(l2.size - 2)
		if err != nil {
			return err
		}
		node.next = nil
		l2.last = node
		l2.size--
		return nil
	}
	// 3. 移除中间的元素
	firstNode, err := l2.node(index - 1)
	if err != nil {
		return err
	}
	tailNode, err := l2.node(index + 1)
	if err != nil {
		return err
	}
	firstNode.next = tailNode
	tailNode.prev = firstNode
	l2.size--
	return nil
}

// 获取index位置对应的节点对象
func (l2 *linkedList) node(index int) (*nodes, error) {
	err := rangeCheck(index, l2.size)
	if err != nil {
		return nil, err
	}
	// 1. > size 的一半，尾部查找
	if index <= (l2.size >> 1) {
		node := l2.first
		for i := 0; i < index; i++ {
			node = node.next
		}
		return node, nil
	} else {
		node := l2.last
		for i := l2.size - 1; i > index; i-- {
			node = node.prev
		}
		return node, nil
	}
}
