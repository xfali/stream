// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package valve

import (
	"math/rand"
	"reflect"
)

type FindFirstValve struct {
	BaseValve
	ret reflect.Value
	set bool
}

func (valve *FindFirstValve) Verify(t reflect.Type) error {
	return nil
}

func (valve *FindFirstValve) Begin(count int) error {
	return nil
}

func (valve *FindFirstValve) End() error {
	return nil
}

func (valve *FindFirstValve) Accept(v reflect.Value) error {
	if !valve.set {
		valve.ret = v
		valve.set = true
	}
	return nil
}

func (valve *FindFirstValve) Result() reflect.Value {
	return valve.ret
}

type FindLastValve struct {
	BaseValve
	ret reflect.Value
}

func (valve *FindLastValve) Verify(t reflect.Type) error {
	return nil
}

func (valve *FindLastValve) Begin(count int) error {
	return nil
}

func (valve *FindLastValve) End() error {
	return nil
}

func (valve *FindLastValve) Accept(v reflect.Value) error {
	valve.ret = v
	return nil
}

func (valve *FindLastValve) Result() reflect.Value {
	return valve.ret
}

type FindAnyValve struct {
	BaseValve
	ret    reflect.Value
	cur    int
	dest   int
	values []reflect.Value
}

func (valve *FindAnyValve) Verify(t reflect.Type) error {
	return nil
}

func (valve *FindAnyValve) Begin(count int) error {
	return nil
}

func (valve *FindAnyValve) End() error {
	return nil
}

func (valve *FindAnyValve) Accept(v reflect.Value) error {
	valve.values = append(valve.values, v)
	return nil
}

func (valve *FindAnyValve) Result() reflect.Value {
	if len(valve.values) == 0 {
		return reflect.Value{}
	} else {
		return valve.values[rand.Intn(len(valve.values))]
	}
}
