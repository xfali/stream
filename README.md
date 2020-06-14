# stream
stream是一个支持多种聚合操作的数据序列，支持转化处理的数据有：
|  类型   | 说明  |
|  :----  | :----  |
| slice  | go切片 |
| list  | go list.List |

支持的方法如下：
|  方法   | 说明  |
|  :----  | :----  |
| Count  | 获得元素的总数 |
| FindFirst  | 获得第一个元素 |
| FindLast  | 获得最后一个元素 |
| FindAny  | 获得一个随机元素 |
| Filter  | 过滤元素，返回一个包括所有符合过滤条件的元素的流 |
| Limit  | 返回一个不超过给定长度的流 |
| Skip  | 返回一个扔掉了前n个元素的流 |
| Distinct  | 返回一个去重的stream |
| Sort  | 返回一个排序后的stream |
| FlatMap  | 映射并扁平化为一个stream |
| Map  | 由一个类型映射到另一个类型 |
| Foreach  | 迭代流中所有数据 |
| Peek  | 获取中间计算的值，主要用于调试 |
| AnyMatch  | 任意匹配一个则返回true，否则返回false |
| AllMatch  | 完全匹配返回true，否则返回false |
| Reduce  | 对stream中元素进行聚合求值 |
| Collect  | 获得slice |

## 惰性求值
stream的实现使用惰性求值（Lazy Evaluation）的计算方式以获取更高的性能。

## 安装
```
go get github.com/xfali/stream
```

## 使用
slice例子：
```
testSlice := []string{"0,1,3,2,2,4", "5,8,7,6,7,9"}
stream.Slice(testSlice).FlatMap(func(s string) []int {
    return stream.Slice(strings.Split(s, ",")).Map(func(s string) int {
        i, _ := strconv.Atoi(s)
        return i
    }).Collect().([]int)
}).Filter(func(i int) bool {
    return i != 5
}).Sort(func(a, b int) int {
    return a - b
}).Distinct(func(a, b int) bool {
    return a == b
}).Map(func(i int) string {
    return strconv.Itoa(i)
}).Foreach(func(s string) {
    t.Log(s)
})
```
list例子：
```
testList := list.New()
testList.PushBack("0,1,3,2,2,4")
testList.PushBack("5,8,7,6,7,9")
stream.List(testList).FlatMap(func(s string) []int {
    return stream.Slice(strings.Split(s, ",")).Map(func(s string) int {
        i, _ := strconv.Atoi(s)
        return i
    }).Collect().([]int)
}).Filter(func(i int) bool {
    return i != 5
}).Sort(func(a, b int) int {
    return a - b
}).Distinct(func(a, b int) bool {
    return a == b
}).Map(func(i int) string {
    return strconv.Itoa(i)
}).Foreach(func(s string) {
    t.Log(s)
})
```

## 未完成项
* 并行处理

