package mylinked

import (
	"fmt"
	"reflect"
)

// 存储的元素
type nodeData struct {
	data interface{}
}

// 单向链表节点
type node struct {
	next *node // 下一个节点指针
	nodeData
}

// 单向链表结构体
type singleLinkedNode struct {
	size  int   // 元素的长度
	first *node // 入口节点
}

// 双向链表节点
type nodes struct {
	prev *nodes
	next *nodes
	nodeData
}

// 双向链表结构体
type linkedNode struct {
	size  int    // 元素的长度
	first *nodes // 头入口节点
	last  *nodes //尾入口节点
}

type SwfengLinkedList interface {
	Clear()                           // 清空链表
	Get(int) interface{}              // 根据索引获取链表的元素
	Set(int, interface{}) interface{} // 根据索引修改元素
	Add(int, ...interface{}) int      // 添加元素
	Remove(int) interface{}           // 根据索引移除元素
	IndexOf(interface{}) int          //根据元素获取索引
	Size() int                        // 获取元素数量
}

const (
	ELEMENT_NOT_FOUND int = -1
)

// 范围检查
func rangeCheck(index, size int) (err error) {
	if index < 0 || index >= size {
		err = fmt.Errorf("Index cross boundary Index:%d , Cap:%d\n", index, size)
	}
	return
}

// ---------------单向链表共用方法------------------

func singleIndexOf(element interface{}, ls *singleLinkedNode) int {
	if element == nil {
		node := ls.first
		for i := 0; i < ls.size; i++ {
			if node.data == nil {
				return i
			}
		}
		node = node.next
	} else {
		node := ls.first
		for i := 0; i < ls.size; i++ {
			if reflect.DeepEqual(element, node.data) {
				return i
			}
			node = node.next
		}
		node = node.next
	}
	return ELEMENT_NOT_FOUND
}
