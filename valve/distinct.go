// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package valve

import "reflect"

type DistinctValve struct {
	BaseValve
	last      reflect.Value
	slice     reflect.Value
	sliceType reflect.Type
}

func (valve *DistinctValve) Verify(t reflect.Type) bool {
	valve.sliceType = reflect.SliceOf(t)
	return VerifyCompareFunction(valve.fn, t)
}

func (valve *DistinctValve) Begin(count int) error {
	if valve.state != DISTINCT && valve.state != SORTED {
		cap := count
		if count == -1 {
			cap = DefaultCapacity
		}
		valve.slice = reflect.MakeSlice(valve.sliceType, 0, cap)
	}
	return valve.next.Begin(count)
}

func (valve *DistinctValve) End() (err error) {
	if valve.state != DISTINCT && valve.state != SORTED {
		for i := 0; i < valve.slice.Len(); i++ {
			err = valve.next.Accept(valve.slice.Index(i))
		}
	}
	valve.state = DISTINCT
	return valve.next.End()
}

func (valve *DistinctValve) Accept(v reflect.Value) error {
	if valve.state == DISTINCT {
		return valve.next.Accept(v)
	} else if valve.state == SORTED {
		if compare(valve.fn, valve.last, v) != 0 {
			valve.last = v
			return valve.next.Accept(v)
		}
	} else {
		found := false
		for i := 0; i < valve.slice.Len(); i++ {
			if compare(valve.fn, valve.slice.Index(i), v) == 0 {
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
