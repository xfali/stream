// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package valve

import (
	"math/rand"
	"reflect"
)

type FindFristValve struct {
	BaseValve
	ret reflect.Value
	set bool
}

func (valve *FindFristValve) Verify(t reflect.Type) bool {
	return true
}

func (valve *FindFristValve) Begin(count int) error {
	return nil
}

func (valve *FindFristValve) End() error {
	return nil
}

func (valve *FindFristValve) Accept(v reflect.Value) error {
	if !valve.set {
		valve.ret = v
		valve.set = true
	}
	return nil
}

func (valve *FindFristValve) Result() reflect.Value {
	return valve.ret
}

type FindLastValve struct {
	BaseValve
	ret reflect.Value
}

func (valve *FindLastValve) Verify(t reflect.Type) bool {
	return true
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
	ret reflect.Value
	cur int
	dest int
	values []reflect.Value
}

func (valve *FindAnyValve) Verify(t reflect.Type) bool {
	return true
}

func (valve *FindAnyValve) Begin(count int) error {
	if count != -1 {
		valve.dest = rand.Intn(count)
	} else {
		valve.dest = -1
	}
	return nil
}

func (valve *FindAnyValve) End() error {
	return nil
}

func (valve *FindAnyValve) Accept(v reflect.Value) error {
	if valve.dest != -1 {
		if valve.cur == valve.dest {
			valve.ret = v
		}
		valve.cur++
	} else {
		valve.values = append(valve.values, v)
	}
	return nil
}

func (valve *FindAnyValve) Result() reflect.Value {
	if valve.dest != -1 {
		return valve.ret
	} else {
		return valve.values[rand.Intn(len(valve.values))]
	}
}
