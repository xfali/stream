// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package collector

import (
	"container/list"
	"errors"
	"github.com/xfali/stream/funcutil"
	"reflect"
)

type Collector interface {
	New(elemType reflect.Type, size int) error
	Add(elemValue reflect.Value) error
	Result() reflect.Value
}

type sliceCollector struct {
	slice reflect.Value
}

func ToSlice() Collector {
	return &sliceCollector{}
}

func (c *sliceCollector) New(elemType reflect.Type, size int) error {
	c.slice = reflect.MakeSlice(reflect.SliceOf(elemType), 0, size)
	return nil
}

func (c *sliceCollector) Add(elemValue reflect.Value) error {
	c.slice = reflect.Append(c.slice, elemValue)
	return nil
}

func (c *sliceCollector) Result() reflect.Value {
	return c.slice
}

type listCollector struct {
	list *list.List
}

func ToList() Collector {
	return &listCollector{}
}

func (c *listCollector) New(elemType reflect.Type, size int) error {
	c.list = list.New()
	return nil
}

func (c *listCollector) Add(elemValue reflect.Value) error {
	c.list.PushBack(elemValue.Interface())
	return nil
}

func (c *listCollector) Result() reflect.Value {
	return reflect.ValueOf(c.list)
}

type mapCollector struct {
	m  reflect.Value
	fn reflect.Value
}

// 转换收集为map
// 参数类型为：fn func(ELEM_TYPE) (KEY_TYPE, VALUE_TYPE)
func ToMap(fn interface{}) Collector {
	return &mapCollector{
		fn: reflect.ValueOf(fn),
	}
}

func (c *mapCollector) New(elemType reflect.Type, size int) error {
	ok := funcutil.VerifyGroupFuncType(c.fn, elemType)
	keyType, ValueType := c.fn.Type().Out(0), c.fn.Type().Out(1)
	if !ok || keyType == nil || ValueType == nil {
		return errors.New("ToMap: Function must be of type func(" + elemType.String() + ") (keyType, valueType)")
	}

	mapType := reflect.MapOf(keyType, ValueType)
	c.m = reflect.MakeMap(mapType)
	return nil
}

func (c *mapCollector) Add(elemValue reflect.Value) error {
	var param [1]reflect.Value
	param[0] = elemValue
	ret := c.fn.Call(param[:])
	c.m.SetMapIndex(ret[0], ret[1])

	return nil
}

func (c *mapCollector) Result() reflect.Value {
	return c.m
}

type groupCollector struct {
	m    reflect.Value
	fn   reflect.Value
	size int
}

// 转换收集为map
// 参数类型为：fn func(ELEM_TYPE) (KEY_TYPE, VALUE_TYPE)
func GroupBy(fn interface{}) Collector {
	return &groupCollector{
		fn: reflect.ValueOf(fn),
	}
}

func (c *groupCollector) New(elemType reflect.Type, size int) error {
	ok := funcutil.VerifyGroupFuncType(c.fn, elemType)
	keyType, ValueType := c.fn.Type().Out(0), c.fn.Type().Out(1)
	if !ok || keyType == nil || ValueType == nil {
		return errors.New("GroupBy: Function must be of type func(" + elemType.String() + ") (keyType, valueType)")
	}

	mapType := reflect.MapOf(keyType, reflect.SliceOf(ValueType))
	c.m = reflect.MakeMap(mapType)
	if size == -1 {
		size = 64
	}
	c.size = size
	return nil
}

func (c *groupCollector) Add(elemValue reflect.Value) error {
	var param [1]reflect.Value
	param[0] = elemValue
	ret := c.fn.Call(param[:])

	v := c.m.MapIndex(ret[0])
	if !v.IsValid() {
		sliceType := c.m.Type().Elem()
		v = reflect.MakeSlice(sliceType, 0, c.size/2)
	}
	v = reflect.Append(v, ret[1])
	c.m.SetMapIndex(ret[0], v)

	return nil
}

func (c *groupCollector) Result() reflect.Value {
	return c.m
}
