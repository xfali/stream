// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package valve

import (
	"reflect"
)

type CollectValve struct {
	BaseValve
	slice reflect.Value
}

func (valve *CollectValve) Verify(t reflect.Type) error {
	valve.slice = reflect.MakeSlice(reflect.SliceOf(t), 0, DefaultCapacity)
	return nil
}

func (valve *CollectValve) Begin(count int) error {
	return nil
}

func (valve *CollectValve) End() error {
	return nil
}

func (valve *CollectValve) Accept(v reflect.Value) error {
	valve.slice = reflect.Append(valve.slice, v)
	return nil
}

func (valve *CollectValve) Result() reflect.Value {
	return valve.slice
}
