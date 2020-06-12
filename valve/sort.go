// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package valve

import (
	"github.com/xfali/stream/funcutil"
	"reflect"
	"sort"
)

type SortValve struct {
	BaseValve
	slice     reflect.Value
	sliceType reflect.Type
}

func (valve *SortValve) Verify(t reflect.Type) bool {
	valve.sliceType = reflect.SliceOf(t)
	return VerifyCompareFunction(valve.fn, t)
}

func (valve *SortValve) Begin(count int) error {
	if !CheckState(valve.state, SORTED) {
		cap := count
		if count == -1 {
			cap = DefaultCapacity
		}
		valve.slice = reflect.MakeSlice(valve.sliceType, 0, cap)
	}
	return valve.next.Begin(count)
}

func (valve *SortValve) End() (err error) {
	if !CheckState(valve.state, SORTED) {
		ss := funcutil.SortSlice{
			V:        valve.slice,
			Compare:  valve.fn,
			SwapFunc: reflect.Swapper(valve.slice.Interface()),
		}
		sort.Sort(&ss)
		for i := 0; i < valve.slice.Len(); i++ {
			err = valve.next.Accept(valve.slice.Index(i))
		}
	}
	valve.state = SetState(valve.state, SORTED)
	return valve.next.End()
}

func (valve *SortValve) Accept(v reflect.Value) error {
	if CheckState(valve.state, SORTED) {
		return valve.next.Accept(v)
	} else {
		valve.slice = reflect.Append(valve.slice, v)
	}
	return nil
}

func (valve *SortValve) Result() reflect.Value {
	return valve.next.Result()
}

func compare(fn, v1, v2 reflect.Value) int64 {
	var param [2]reflect.Value
	param[0] = v1
	param[1] = v2
	return fn.Call(param[:])[0].Int()
}
