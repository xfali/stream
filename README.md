# stream
stream是一个数据处理工具，支持方法如下：
|  方法   | 说明  |
|  :----  | :----  |
| Count  | 获得slice的总数 |
| FindFirst  | 获得第一个元素 |
| FindAny  | 获得一个随机元素 |
| FindFirst  | 单元格 |
| Filter  | 过滤元素 |
| Map  | 由一个类型映射到另一个类型 |
| Foreach  | 迭代流中所有数据 |
| AnyMatch  | 任意匹配一个则返回true，否则返回false |
| AllMatch  | 完全匹配返回true，否则返回false |
| Reduce  | 对stream中元素进行聚合求值 |
| Collect  | 获得slice |

## 安装
```cassandraql
go get github.com/xfali/stream
```

## 使用
例子：
```cassandraql
s := stream.Slice([]int{1, 2, 3, 4, 5})
s.Filter(func(i int) bool {
    if i == 2 {
        return false
    }
    return true
}).Foreach(func(i int) {
    if i == 2 {
        t.Fatal("filter 2 but got it")
    }
})
```

## 未完成项
* 持续丰富stream API
* 提高性能
* 并行处理

