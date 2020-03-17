## 链表工具库
- 已实现单向链表，双向链表，单向循环链表，双向循环链表

### 接口
- Clear()   // 清空元素
- Add(index int, args ...interface{}) error // 添加元素
- Remove(index int) (interface{}, error) // 删除元素
- IndexOf(E element) int   // 获取元素的索引
- Get(index int) (interface{}, error)    //获取元素
- Size() int    // 获取链表容量大小
- Set(index int, element interface{}) (interface{}, error)  //设置元素


### 基准测试：
- 单向链表
```
$  go test -bench=. -benchmem
goos: windows
goarch: amd64
pkg: shiwifeng/linkedlist
BenchmarkSingleLinkedList_Add-4            10000            318318 ns/op             128 B/op          4 allocs/op
BenchmarkSingleLinkedList_Get-4             6670            194494 ns/op               0 B/op          0 allocs/op
BenchmarkSingleLinkedList_IndexOf-4       107204             12231 ns/op               0 B/op          0 allocs/op
BenchmarkSingleLinkedList_Remove-4         18055             91723 ns/op               0 B/op          0 allocs/op
BenchmarkSingleLinkedList_Set-4            21830             57457 ns/op               0 B/op          0 allocs/op
PASS
ok      shiwifeng/linkedlist    52.955s

```
```
package linkedlist

import (
	"math/rand"
	"testing"
	"time"
)

var dd2 *SingleLinkedList

func BenchmarkSingleLinkedList_Add(b *testing.B) {
	dd2 = NewSingleLinkedList()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < b.N; i++ {
		if dd2.Size() < 1 {
			dd2.Add(0, data...)
		} else {
			rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	dd2.Clear()
}

func BenchmarkSingleLinkedList_Get(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < 10000; i++ {
		if dd2.Size() < 1 {
			dd2.Add(0, data...)
		} else {
			rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
		dd2.Get(rand.Intn(dd2.Size()))
	}
	dd2.Clear()
}

func BenchmarkSingleLinkedList_IndexOf(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < 10000; i++ {
		if dd2.Size() < 1 {
			dd2.Add(0, data...)
		} else {
			rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
		dd2.IndexOf(data[rand.Intn(4)])
	}
	dd2.Clear()
}

func BenchmarkSingleLinkedList_Remove(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < 10000; i++ {
		if dd2.Size() < 1 {
			dd2.Add(dd2.Size(), data...)
		} else {
			rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
		dd2.Remove(rand.Intn(10000))
	}
	dd2.Clear()
}

func BenchmarkSingleLinkedList_Set(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < 10000; i++ {
		if dd2.Size() < 1 {
			dd2.Add(dd2.Size(), data...)
		} else {
			rand.Seed(time.Now().UnixNano()+int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano()+int64(i)) //随机种子
		dd2.Set(rand.Intn(10000),9999)
	}
	dd2.Clear()
}

```
- 双向链表
```
$  go test -bench=. -benchmem
goos: windows
goarch: amd64
pkg: shiwifeng/linkedlist
BenchmarkLinkedList_Add-4                  20316            117722 ns/op             128 B/op          4 allocs/op
BenchmarkLinkedList_Get-4                  13136             91680 ns/op               0 B/op          0 allocs/op
BenchmarkLinkedList_IndexOf-4              99229             12006 ns/op               0 B/op          0 allocs/op
BenchmarkLinkedList_Remove-4               21751             61525 ns/op               0 B/op          0 allocs/op
BenchmarkLinkedList_Set-4                  25170             44353 ns/op               0 B/op          0 allocs/op
PASS
ok      shiwifeng/linkedlist    19.450s

```
```
package linkedlist

import (
	"math/rand"
	"testing"
	"time"
)

var dd2 *LinkedList

func BenchmarkLinkedList_Add(b *testing.B) {
	dd2 = NewLinkedList()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < b.N; i++ {
		if dd2.Size() < 1 {
			dd2.Add(0, data...)
		} else {
			rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	dd2.Clear()
}

func BenchmarkLinkedList_Get(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < 10000; i++ {
		if dd2.Size() < 1 {
			dd2.Add(0, data...)
		} else {
			rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
		dd2.Get(rand.Intn(dd2.Size()))
	}
	dd2.Clear()
}

func BenchmarkLinkedList_IndexOf(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < 10000; i++ {
		if dd2.Size() < 1 {
			dd2.Add(0, data...)
		} else {
			rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
		dd2.IndexOf(data[rand.Intn(4)])
	}
	dd2.Clear()
}

func BenchmarkLinkedList_Remove(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < 10000; i++ {
		if dd2.Size() < 1 {
			dd2.Add(dd2.Size(), data...)
		} else {
			rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
		dd2.Remove(rand.Intn(10000))
	}
	dd2.Clear()
}

func BenchmarkLinkedList_Set(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < 10000; i++ {
		if dd2.Size() < 1 {
			dd2.Add(dd2.Size(), data...)
		} else {
			rand.Seed(time.Now().UnixNano()+int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano()+int64(i)) //随机种子
		dd2.Set(rand.Intn(10000),9999)
	}
	dd2.Clear()
}

```
- 单向闭环链表
```
$  go test -bench=. -benchmem
goos: windows
goarch: amd64
pkg: shiwifeng/linkedlist
BenchmarkSingleCircleLinkedList_Add-4              10000            273344 ns/op             128 B/op          4 allocs/op
BenchmarkSingleCircleLinkedList_Get-4               6670            191944 ns/op               0 B/op          0 allocs/op
BenchmarkSingleCircleLinkedList_IndexOf-4         107185             11310 ns/op               0 B/op          0 allocs/op
BenchmarkSingleCircleLinkedList_Remove-4           21751             59458 ns/op               0 B/op          0 allocs/op
BenchmarkSingleCircleLinkedList_Set-4              23728             49913 ns/op               0 B/op          0 allocs/op
PASS
ok      shiwifeng/linkedlist    46.967s

```
```
package linkedlist

import (
	"math/rand"
	"testing"
	"time"
)

var dd2 *SingleCircleLinkedList


func BenchmarkSingleCircleLinkedList_Add(b *testing.B) {
	dd2 = NewSingleCircleLinkedList()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < b.N; i++ {
		if dd2.Size() < 1 {
			dd2.Add(0, data...)
		} else {
			rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	dd2.Clear()
}

func BenchmarkSingleCircleLinkedList_Get(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < 10000; i++ {
		if dd2.Size() < 1 {
			dd2.Add(0, data...)
		} else {
			rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
		dd2.Get(rand.Intn(dd2.Size()))
	}
	dd2.Clear()
}

func BenchmarkSingleCircleLinkedList_IndexOf(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < 10000; i++ {
		if dd2.Size() < 1 {
			dd2.Add(0, data...)
		} else {
			rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
		dd2.IndexOf(data[rand.Intn(4)])
	}
	dd2.Clear()
}

func BenchmarkSingleCircleLinkedList_Remove(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < 10000; i++ {
		if dd2.Size() < 1 {
			dd2.Add(dd2.Size(), data...)
		} else {
			rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
		dd2.Remove(rand.Intn(10000))
	}
	dd2.Clear()
}

func BenchmarkSingleCircleLinkedList_Set(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < 10000; i++ {
		if dd2.Size() < 1 {
			dd2.Add(dd2.Size(), data...)
		} else {
			rand.Seed(time.Now().UnixNano()+int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano()+int64(i)) //随机种子
		dd2.Set(rand.Intn(10000),9999)
	}
	dd2.Clear()
}

```

- 双向闭环链表
```
$  go test -bench=. -benchmem
goos: windows
goarch: amd64
pkg: shiwifeng/linkedlist
BenchmarkCircleLinkedList_Add-4            24157            122544 ns/op             128 B/op          4 allocs/op
BenchmarkCircleLinkedList_Get-4            14577             90845 ns/op               0 B/op          0 allocs/op
BenchmarkCircleLinkedList_IndexOf-4       106254             11184 ns/op               0 B/op          0 allocs/op
BenchmarkCircleLinkedList_Remove-4         22234             60008 ns/op               0 B/op          0 allocs/op
BenchmarkCircleLinkedList_Set-4            25276             53656 ns/op               0 B/op          0 allocs/op
PASS
ok      shiwifeng/linkedlist    20.844s

```
```
package linkedlist

import (
	"math/rand"
	"testing"
	"time"
)

var dd2 *CircleLinkedList

func BenchmarkCircleLinkedList_Add(b *testing.B) {
	dd2 = NewCircleLinkedList()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < b.N; i++ {
		if dd2.Size() < 1 {
			dd2.Add(0, data...)
		} else {
			rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	dd2.Clear()
}

func BenchmarkCircleLinkedList_Get(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < 10000; i++ {
		if dd2.Size() < 1 {
			dd2.Add(0, data...)
		} else {
			rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
		dd2.Get(rand.Intn(dd2.Size()))
	}
	dd2.Clear()
}

func BenchmarkCircleLinkedList_IndexOf(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < 10000; i++ {
		if dd2.Size() < 1 {
			dd2.Add(0, data...)
		} else {
			rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
		dd2.IndexOf(data[rand.Intn(4)])
	}
	dd2.Clear()
}

func BenchmarkCircleLinkedList_Remove(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < 10000; i++ {
		if dd2.Size() < 1 {
			dd2.Add(dd2.Size(), data...)
		} else {
			rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i)) //随机种子
		dd2.Remove(rand.Intn(10000))
	}
	dd2.Clear()
}

func BenchmarkCircleLinkedList_Set(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	s1 := []int{1, 2, 3}
	m1 := make(map[int]int, 3)
	m1[11] = 11
	m1[22] = 22
	m1[33] = 33
	data := []interface{}{
		s1,
		m1,
		99,
		"加油",
	}
	for i := 0; i < 10000; i++ {
		if dd2.Size() < 1 {
			dd2.Add(dd2.Size(), data...)
		} else {
			rand.Seed(time.Now().UnixNano()+int64(i)) //随机种子
			dd2.Add(rand.Intn(dd2.Size()), data...)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano()+int64(i)) //随机种子
		dd2.Set(rand.Intn(10000),9999)
	}
	dd2.Clear()
}

```


