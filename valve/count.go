// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package valve

import (
	"reflect"
)

type CountValve struct {
	BaseValve
	cur   int
	count int
}

func (valve *CountValve) Verify(t reflect.Type) bool {
	return true
}

func (valve *CountValve) Begin(count int) error {
	valve.count = count
	return nil
}

func (valve *CountValve) End() error {
	return nil
}

func (valve *CountValve) Accept(v reflect.Value) error {
	if valve.count == -1 {
		valve.cur++
	}
	return nil
}

func (valve *CountValve) Result() reflect.Value {
	if valve.count != -1 {
		valve.cur = valve.count
	}
	return reflect.ValueOf(valve.cur)
}
