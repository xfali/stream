// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package valve

import (
	"reflect"
)

type NoneValve struct {
	BaseValve
}

func (valve *NoneValve) Verify(t reflect.Type) error {
	return valve.next.Verify(t)
}

func (valve *NoneValve) Begin(count int) error {
	return valve.next.Begin(count)
}

func (valve *NoneValve) End() error {
	return valve.next.End()
}

func (valve *NoneValve) Accept(v reflect.Value) error {
	return valve.next.Accept(v)
}
func (valve *NoneValve) Result() reflect.Value {
	return valve.next.Result()
}
