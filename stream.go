// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package stream

type Stream interface {
	// 获得元素的总数
	Count() int

	// 获得第一个元素
	FindFirst() *Option

	// 获得最后一个元素
	// Deprecated: 在使用pipeline方式时间复杂度为O(n)
	FindLast() *Option

	// 获得一个随机元素
	FindAny() *Option

	// 过滤元素，返回一个包括所有符合过滤条件的元素的stream
	// 参数类型为：fn func(TYPE) bool
	Filter(fn interface{}) Stream

	// 返回一个不超过给定长度的stream
	Limit(size int) Stream

	// 返回一个扔掉了前n个元素的stream
	Skip(size int) Stream

	// 返回一个去重的stream
	// 参数类型为：fn func(t1, t2 TYPE) bool
	Distinct(fn interface{}) Stream

	// 返回一个排序后的stream
	// 参数类型为：fn func(t1, t2 TYPE) int
	Sort(fn interface{}) Stream

	// 映射并扁平化为一个stream
	// 参数类型为：fn func(OLD_TYPE) []NEW_TYPE
	FlatMap(fn interface{}) Stream

	// 由一个类型映射到另一个类型
	// 参数类型为：fn func(OLD_TYPE) NEW_TYPE
	Map(fn interface{}) Stream

	// 迭代流中所有数据
	// 参数类型为：fn func(TYPE)
	Foreach(fn interface{})

	// 迭代流中所有数据，并返回stream
	// 参数类型为：fn func(TYPE)
	Peek(fn interface{}) Stream

	// 任意匹配一个则返回true，否则返回false
	// 参数类型为：fn func(TYPE) bool
	AnyMatch(fn interface{}) bool

	// 完全匹配返回true，否则返回false
	// 参数类型为：fn func(o TYPE) bool
	AllMatch(fn interface{}) bool

	// 对stream中元素进行聚合求值
	// 参数类型为：fn func(out, in TYPE) interface{}
	Reduce(fn, initValue interface{}) interface{}

	// 获得slice
	// 参数类型为：collector.Collector
	Collect(collector interface{}) interface{}
}
