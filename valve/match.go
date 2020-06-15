// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package valve

import "reflect"

type MatchAnyValve struct {
	BaseValve
	ret bool
}

func (valve *MatchAnyValve) Reset() {
	valve.ret = false
}

func (valve *MatchAnyValve) Verify(t reflect.Type) error {
	return VerifyFilterFuncType(valve.fn, t)
}

func (valve *MatchAnyValve) Begin(count int) error {
	return nil
}

func (valve *MatchAnyValve) End() error {
	return nil
}

func (valve *MatchAnyValve) Accept(v reflect.Value) error {
	if !valve.ret {
		var param [1]reflect.Value
		param[0] = v
		if valve.fn.Call(param[:])[0].Bool() {
			valve.ret = true
		}
	}
	return nil
}

func (valve *MatchAnyValve) Result() reflect.Value {
	return reflect.ValueOf(valve.ret)
}

type MatchAllValve struct {
	BaseValve
	ret bool
}

func (valve *MatchAllValve) Reset() {
	valve.ret = false
}

func (valve *MatchAllValve) Verify(t reflect.Type) error {
	return VerifyFilterFuncType(valve.fn, t)
}

func (valve *MatchAllValve) Begin(count int) error {
	return nil
}

func (valve *MatchAllValve) End() error {
	return nil
}

func (valve *MatchAllValve) Accept(v reflect.Value) error {
	if !valve.ret {
		var param [1]reflect.Value
		param[0] = v
		if !valve.fn.Call(param[:])[0].Bool() {
			valve.ret = true
		}
	}
	return nil
}

func (valve *MatchAllValve) Result() reflect.Value {
	return reflect.ValueOf(!valve.ret)
}
