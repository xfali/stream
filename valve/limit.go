// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package valve

import (
	"reflect"
)

type LimitValve struct {
	BaseValve
	ret   reflect.Value
	Limit int
	cur   int
}

func (valve *LimitValve) Verify(t reflect.Type) bool {
	return true
}

func (valve *LimitValve) Begin(count int) error {
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

func (valve *LimitValve) End() error {
	return nil
}

func (valve *LimitValve) Accept(v reflect.Value) error {
	if valve.cur < valve.Limit {
		return valve.next.Accept(v)
	}
	return nil
}

func (valve *LimitValve) Result() reflect.Value {
	return valve.next.Result()
}

type SkipValve struct {
	BaseValve
	ret  reflect.Value
	Skip int
	cur  int
}

func (valve *SkipValve) Verify(t reflect.Type) bool {
	return true
}

func (valve *SkipValve) Begin(count int) error {
	if count == -1 {
		return valve.next.Begin(-1)
	} else {
		if valve.Skip > count {
			return valve.next.Begin(0)
		} else {
			return valve.next.Begin(count - valve.Skip)
		}
	}
}

func (valve *SkipValve) End() error {
	return nil
}

func (valve *SkipValve) Accept(v reflect.Value) error {
	if valve.cur >= valve.Skip {
		return valve.next.Accept(v)
	}
	return nil
}

func (valve *SkipValve) Result() reflect.Value {
	return valve.next.Result()
}
