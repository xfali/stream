// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package funcutil

import (
	"errors"
	"reflect"
	sort2 "sort"
)

func Sort(function, slice interface{}) (ret interface{}, err error) {
	return sort(function, slice)
}

func sort(function, slice interface{}) (interface{}, error) {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		return nil, errors.New("Limit The first param is not a slice ")
	}
	fn := reflect.ValueOf(function)
	inType := in.Type().Elem()
	if !VerifyCompareFunction(fn, inType) {
		panic("sort: Function must be of type func(" + inType.String() + "," + inType.String() + ") int")
	}

	ss := SortSlice{
		V:       in,
		Compare: fn,
		SwapFunc:    reflect.Swapper(slice),
	}
	sort2.Sort(&ss)
	return ss.V.Interface(), nil
}

type SortSlice struct {
	V        reflect.Value
	Compare  reflect.Value
	SwapFunc func(i, j int)
}

func (p *SortSlice) Len() int {
	return p.V.Len()
}

func (p *SortSlice) Less(i, j int) bool {
	var param [2]reflect.Value
	param[0] = p.V.Index(i)
	param[1] = p.V.Index(j)
	return p.Compare.Call(param[:])[0].Int() < 0
}

func (p *SortSlice) Swap(i, j int) {
	p.SwapFunc(i, j)
}

func VerifyCompareFunction(fn reflect.Value, elemType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 2 || fn.Type().NumOut() != 1 {
		return false
	}
	if elemType != fn.Type().In(0) || elemType != fn.Type().In(1) || reflect.Int != fn.Type().Out(0).Kind() {
		return false
	}
	return true
}