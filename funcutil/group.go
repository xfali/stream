// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package funcutil

import (
	"errors"
	"reflect"
)

func Group(function, slice interface{}) (ret interface{}, err error) {
	return group(function, slice)
}

func group(function, slice interface{}) (interface{}, error) {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		return nil, errors.New("Limit The first param is not a slice ")
	}
	fn := reflect.ValueOf(function)
	inType := in.Type().Elem()
	ok := VerifyGroupFuncType(fn, inType)
	keyType, ValueType := fn.Type().Out(0), fn.Type().Out(1)
	if !ok || keyType == nil || ValueType == nil {
		panic("group: Function must be of type func(" + inType.String() + ") (keyType, valueType)")
	}

	sliceType := reflect.SliceOf(ValueType)
	mapType := reflect.MapOf(keyType, sliceType)
	out := reflect.MakeMap(mapType)
	var param [1]reflect.Value
	for i := 0; i < in.Len(); i++ {
		param[0] = in.Index(i)
		ret := fn.Call(param[:])
		v := out.MapIndex(ret[0])
		if !v.IsValid() {
			v = reflect.MakeSlice(sliceType, 0, in.Len()/2)
		}
		v = reflect.Append(v, ret[1])
		out.SetMapIndex(ret[0], v)
	}

	return out.Interface(), nil
}

func VerifyGroupFuncType(fn reflect.Value, elemType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 2 {
		return false
	}
	if fn.Type().In(0) != elemType {
		return false
	}
	return true
}
