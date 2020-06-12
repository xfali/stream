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
	ret   reflect.Value
	Limit int
	cur   int
}

func (valve *CountValve) Verify(t reflect.Type) bool {
	return true
}

func (valve *CountValve) Begin(count int) error {
	if count == -1 {
		return valve.next.Begin(-1)
	} else {
		if valve.Limit > count {
			return valve.next.Begin(count)
		} else {
			return valve.next.Begin(valve.Limit)
		}
	}
}

func (valve *CountValve) End() error {
	return nil
}

func (valve *CountValve) Accept(v reflect.Value) error {
	if valve.cur < valve.Limit {
		valve.cur++
		return valve.next.Accept(v)
	}
	valve.cur++
	return nil
}

func (valve *CountValve) Result() reflect.Value {
	return valve.next.Result()
}
