// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package funcutil

import (
	"errors"
	"reflect"
)

func Distinct(function, slice interface{}) (ret interface{}, err error) {
	return distinct(function, slice)
}

func distinct(function, slice interface{}) (interface{}, error) {
	in := reflect.ValueOf(slice)
	fn := reflect.ValueOf(function)
	if in.Kind() != reflect.Slice {
		return nil, errors.New("Limit The first param is not a slice ")
	}
	inType := in.Type().Elem()
	if !VerifyEqualFunction(fn, inType) {
		panic("distinct: Function must be of type func(" + inType.String() + "," + inType.String() + ") int")
	}

	out := reflect.MakeSlice(in.Type(), 0, in.Len())
	out = reflect.Append(out, in.Index(0))
	var param [2]reflect.Value
	for i := 1; i < in.Len(); i++ {
		found := false
		for j := 0; j < out.Len(); j++ {
			param[0] = in.Index(i)
			param[1] = out.Index(j)
			if fn.Call(param[:])[0].Bool() {
				found = true
			}
		}
		if !found {
			out = reflect.Append(out, in.Index(i))
		}
	}
	return out.Interface(), nil
}

func VerifyEqualFunction(fn reflect.Value, elemType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 2 || fn.Type().NumOut() != 1 {
		return false
	}
	if elemType != fn.Type().In(0) || elemType != fn.Type().In(1) || reflect.Bool != fn.Type().Out(0).Kind() {
		return false
	}
	return true
}
