// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package valve

import "reflect"

type FilterValve struct {
	BaseValve
}

func (valve *FilterValve) Verify(t reflect.Type) bool {
	return VerifyFilterFuncType(valve.fn, t) && valve.next.Verify(t)
}

func (valve *FilterValve) Begin(count int) error {
	return valve.next.Begin(-1)
}

func (opt *FilterValve) End() error {
	return opt.next.End()
}

func (valve *FilterValve) Accept(v reflect.Value) error {
	var param [1]reflect.Value
	param[0] = v
	if valve.fn.Call(param[:])[0].Bool() {
		return valve.next.Accept(v)
	}
	return nil
}

func (valve *FilterValve) Result() reflect.Value {
	return valve.next.Result()
}

func VerifyFilterFuncType(fn reflect.Value, elemType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
		return false
	}
	if fn.Type().In(0) != elemType || fn.Type().Out(0).Kind() != reflect.Bool {
		return false
	}
	return true
}
