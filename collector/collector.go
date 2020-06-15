// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package collector

import (
	"container/list"
	"reflect"
)

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
