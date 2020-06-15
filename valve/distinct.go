// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package valve

import (
	"fmt"
	"github.com/xfali/stream/funcutil"
	"reflect"
)

type DistinctValve struct {
	BaseValve
	last      reflect.Value
	slice     reflect.Value
	sliceType reflect.Type
}

func (valve *DistinctValve) Reset() {
	valve.last = reflect.Value{}
	valve.slice = reflect.Value{}
	valve.next.Reset()
}

func (valve *DistinctValve) Verify(t reflect.Type) error {
	valve.sliceType = reflect.SliceOf(t)
	if !funcutil.VerifyEqualFunction(valve.fn, t) {
		return fmt.Errorf("distinct: Function must be of type func(%s, %s) bool", t.String(), t.String())
	}
	return valve.next.Verify(t)
}

func (valve *DistinctValve) Begin(count int) error {
	if !CheckState(valve.state, DISTINCT) && !CheckState(valve.state, SORTED) {
		cap := count
		if count == -1 {
			cap = DefaultCapacity
		}
		valve.slice = reflect.MakeSlice(valve.sliceType, 0, cap)
	}
	valve.next.SetState(SetState(valve.state, DISTINCT))
	return valve.next.Begin(count)
}

func (valve *DistinctValve) End() (err error) {
	if !CheckState(valve.state, DISTINCT) && !CheckState(valve.state, SORTED) {
		for i := 0; i < valve.slice.Len(); i++ {
			err = valve.next.Accept(valve.slice.Index(i))
		}
	}
	valve.state = SetState(valve.state, DISTINCT)
	return valve.next.End()
}

func (valve *DistinctValve) Accept(v reflect.Value) error {
	if CheckState(valve.state, DISTINCT) {
		return valve.next.Accept(v)
	} else if CheckState(valve.state, SORTED) {
		if !valve.last.IsValid() || !equal(valve.fn, valve.last, v) {
			valve.last = v
			return valve.next.Accept(v)
		}
	} else {
		found := false
		for i := 0; i < valve.slice.Len(); i++ {
			if equal(valve.fn, valve.slice.Index(i), v) {
				found = true
			}
		}
		if !found {
			valve.slice = reflect.Append(valve.slice, v)
		}
	}
	return nil
}

func (valve *DistinctValve) Result() reflect.Value {
	return valve.next.Result()
}

func equal(fn, v1, v2 reflect.Value) bool {
	var param [2]reflect.Value
	param[0] = v1
	param[1] = v2
	return fn.Call(param[:])[0].Bool()
}
