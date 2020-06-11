// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package funcutil

import (
	"errors"
	"reflect"
)

// Apply applies a function(the first parameter) to each element of a slice(second parameter). It acts just like Map in other languages.
func Map(function, slice interface{}) (ret interface{}, err error) {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		return nil, errors.New("map: The first param is not a slice ")
	}
	fn := reflect.ValueOf(function)
	inType := in.Type().Elem()
	ok, outType := verifyMapFuncType(fn, inType)
	if !ok {
		return nil, errors.New("map: Function must be of type func(" + inType.String() + ") outputElemType")
	}
	var param [1]reflect.Value
	out := reflect.MakeSlice(reflect.SliceOf(outType), 0, in.Len())
	for i := 0; i < in.Len(); i++ {
		param[0] = in.Index(i)
		out = reflect.Append(out, fn.Call(param[:])[0])
	}
	return out.Interface(), nil
}

func verifyMapFuncType(fn reflect.Value, elemType reflect.Type) (bool, reflect.Type) {
	if fn.Kind() != reflect.Func {
		return false, nil
	}
	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
		return false, nil
	}
	if fn.Type().In(0) != elemType {
		return false, nil
	}
	return true, fn.Type().Out(0)
}
