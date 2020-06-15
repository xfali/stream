// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package collector

import (
	"container/list"
	"reflect"
)

type CollectFunc func(elemType reflect.Type) func(int) func(elemValue reflect.Value) reflect.Value

func CollectSliceFunc(elemType reflect.Type) func(int) func(elemValue reflect.Value) reflect.Value {
	return func(count int) func(elemValue reflect.Value) reflect.Value {
		v := reflect.MakeSlice(reflect.SliceOf(elemType), 0, count)
		return func(elemValue reflect.Value) reflect.Value {
			v = reflect.Append(v, elemValue)
			return v
		}
	}
}

func CollectListFunc(elemType reflect.Type) func(int) func(elemValue reflect.Value) reflect.Value {
	return func(count int) func(elemValue reflect.Value) reflect.Value {
		v := list.New()
		value := reflect.ValueOf(v)
		return func(elemValue reflect.Value) reflect.Value {
			v.PushBack(elemValue.Interface())
			return value
		}
	}
}

func Slice(elemType reflect.Type) func(value reflect.Value) reflect.Value {
	t := reflect.SliceOf(elemType)
	v := reflect.MakeSlice(t, 0, 1024)
	return func(value reflect.Value) reflect.Value {
		v = reflect.Append(v, value)
		return v
	}
}

func List() func(value reflect.Value) reflect.Value {
	t := list.New()
	return func(value reflect.Value) reflect.Value {
		t.PushBack(value.Interface())
		return reflect.ValueOf(t)
	}
}
