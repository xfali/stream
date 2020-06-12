// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package valve

import (
	"github.com/xfali/stream/funcutil"
	"reflect"
)

type MapValve struct {
	BaseValve
}

func (valve *MapValve) Verify(t reflect.Type) bool {
	if !funcutil.VerifyMapFuncType(valve.fn, t) {
		return false
	}
	return valve.next.Verify(valve.fn.Type().Out(0))
}

func (valve *MapValve) Begin(count int) error {
	return valve.next.Begin(count)
}

func (valve *MapValve) End() error {
	return nil
}

func (valve *MapValve) Accept(v reflect.Value) error {
	var param [1]reflect.Value
	param[0] = v
	return valve.next.Accept(valve.fn.Call(param[:])[0])
}

func (valve *MapValve) Result() reflect.Value {
	return valve.next.Result()
}