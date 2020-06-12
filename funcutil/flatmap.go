// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package funcutil

import (
	"errors"
	"reflect"
)

func FlatMap(function, slice interface{}) (interface{}, error) {
	in := reflect.ValueOf(slice)
	fn := reflect.ValueOf(function)
	if in.Kind() != reflect.Slice {
		return nil, errors.New("Limit The first param is not a slice ")
	}
	inType := in.Type().Elem()
	ok := VerifyFlatMapFunction(fn, inType)
	outType := fn.Type().Out(0)
	if !ok {
		panic("flatmap: Function must be of type func(" + inType.String() + ") []interface{}")
	}

	out := reflect.MakeSlice(outType, 0, in.Len())
	var param [1]reflect.Value
	for i := 0; i < in.Len(); i++ {
		param[0] = in.Index(i)
		newSlice := fn.Call(param[:])[0]
		for j := 0; j < newSlice.Len(); j++ {
			out = reflect.Append(out, newSlice.Index(j))
		}
	}
	return out.Interface(), nil
}

func VerifyFlatMapFunction(fn reflect.Value, elemType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
		return false
	}
	if elemType != fn.Type().In(0) || reflect.Slice != fn.Type().Out(0).Kind() {
		return false
	}
	return true
}
