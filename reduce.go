// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package stream

import (
	"errors"
	"reflect"
)

// Reduce applies a function (the first parameter) of two arguments cumulatively to each element of a slice (second parameter), and the initial value is the third parameter.
func Reduce(function, slice interface{}) (ret interface{}, err error) {
	return reduce(function, slice)
}

func reduce(function, slice interface{}) (interface{}, error) {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		return nil, errors.New("Reduce: The first param is not a slice ")
	}

	if in.Len() == 0 {
		return nil, nil
	}

	if in.Len() == 1 {
		return in.Index(0).Interface(), nil
	}
	fn := reflect.ValueOf(function)
	inType := in.Type().Elem()
	if !verifyReduceFuncType(fn, inType) {
		panic("reduce: Function must be of type func(" + inType.String() + ")" + inType.String())
	}

	var param [2]reflect.Value
	out := in.Index(0)
	for i := 1; i < in.Len(); i++ {
		param[0] = out
		param[1] = in.Index(i)
		out = fn.Call(param[:])[0]
	}
	return out.Interface(), nil
}

func verifyReduceFuncType(fn reflect.Value, elemType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 2 || fn.Type().NumOut() != 1 {
		return false
	}
	if elemType != fn.Type().In(0) || fn.Type().In(0) != fn.Type().In(1) || fn.Type().In(1) != fn.Type().Out(0) {
		return false
	}
	return true
}
