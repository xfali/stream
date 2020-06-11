// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package valve

import "reflect"

type ForeachValve struct {
	BaseValve
}

func (valve *ForeachValve) Verify(t reflect.Type) bool {
	return VerifyForeachFuncType(valve.fn, t)
}

func (valve *ForeachValve) Begin(count int) error {
	return nil
}

func (valve *ForeachValve) End() error {
	return nil
}

func (valve *ForeachValve) Accept(v reflect.Value) error {
	var param [1]reflect.Value
	param[0] = v
	valve.fn.Call(param[:])
	return nil
}

func (valve *ForeachValve) Result() reflect.Value {
	return reflect.Value{}
}

func VerifyForeachFuncType(fn reflect.Value, elemType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 0 {
		return false
	}
	if fn.Type().In(0) != elemType {
		return false
	}
	return true
}
