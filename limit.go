// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package stream

import (
	"errors"
	"reflect"
)

func Limit(size int, slice interface{}) (ret interface{}, err error) {
	return limit(size, slice)
}

func limit(size int, slice interface{}) (interface{}, error) {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		return nil, errors.New("Limit The first param is not a slice ")
	}
	if size >= in.Len() {
		return slice, nil
	}
	out := reflect.MakeSlice(in.Type(), 0, size)
	for i := 0; i < size; i++ {
		out = reflect.Append(out, in.Index(i))
	}
	return out.Interface(), nil
}

func Skip(size int, slice interface{}) (ret interface{}, err error) {
	return skip(size, slice)
}

func skip(size int, slice interface{}) (interface{}, error) {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		return nil, errors.New("Limit The first param is not a slice ")
	}
	if size == 0 {
		return slice, nil
	}
	if size > in.Len() {
		size = in.Len()
	}
	outSize := in.Len() - size
	out := reflect.MakeSlice(in.Type(), 0, outSize)
	for i := size; i < in.Len(); i++ {
		out = reflect.Append(out, in.Index(i))
	}
	return out.Interface(), nil
}
