## 链表工具库
- 以实现单向链表，双向链表，单向循环链表，双向循环链表

### 接口
- Clear()   // 清空元素
- Add(index int, args ...interface{}) error // 添加元素
- Remove(index int) (interface{}, error) // 删除元素
- IndexOf(E element) int   // 获取元素的索引
- Get(index int) (interface{}, error)    //获取元素
- Size() int    // 获取链表容量大小
- Set(index int, element interface{}) (interface{}, error)  //设置元素


### 用法示例：
```


func main() {
	dd1 := NewSingleLinkedList()    //单向链表实例
        // dd2 := NewLinkedList() //双向链表示例
        // dd3 := NewSingleCircleLinkedList() //单向循环链表示例
        // dd4 := NewCircleLinkedList() //双向循环链表示例
	// 1. 第一次添加
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	ret := dd1.Add(0, 1, 2, 3, s1, m1)
	if ret != nil {
		// TODO
	}
	// 2. 尾部追加
	ret = dd1.Add(dd1.Size(), 66, 77)
	if ret != nil {
		// TODO
	}
	// 3. 中间插入
	ret = dd1.Add(dd1.Size()>>1, 44, 55)
	if ret != nil {
		// TODO
	}
	// 4. 头部插入
	ret = dd1.Add(0, 22, 33)
	if ret != nil {
		// TODO
	}
    
	ret, _ := dd1.Set(1, 99)
	ret2 ,_ := dd1.Get(1)
	if !reflect.DeepEqual(ret,ret2) {
		// TODO
	}
}
```