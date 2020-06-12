// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package valve

import (
	"reflect"
)

type PeekValve struct {
	BaseValve
}

func (valve *PeekValve) Verify(t reflect.Type) error {
	if err := VerifyForeachFuncType(valve.fn, t); err != nil {
		return err
	}
	if valve.next != nil {
		return valve.next.Verify(t)
	}
	return nil
}

func (valve *PeekValve) Begin(count int) error {
	if valve.next != nil {
		return valve.next.Begin(count)
	}
	return nil
}

func (valve *PeekValve) End() error {
	if valve.next != nil {
		return valve.next.End()
	}
	return nil
}

func (valve *PeekValve) Accept(v reflect.Value) error {
	var param [1]reflect.Value
	param[0] = v
	valve.fn.Call(param[:])
	if valve.next != nil {
		return valve.next.Accept(v)
	}
	return nil
}

func (valve *PeekValve) Result() reflect.Value {
	if valve.next != nil {
		return valve.next.Result()
	}
	return reflect.Value{}
}
