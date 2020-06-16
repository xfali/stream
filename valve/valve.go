// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package valve

import (
	"errors"
	"reflect"
)

const (
	NORMAL   = 0
	DISTINCT = 1
	SORTED   = 1 << 1
)

var DefaultCapacity = 64

type FuncValve interface {
	Init(fn interface{}) error
	Reset()
	Next(v FuncValve)

	Verify(t reflect.Type) error

	Begin(t int) error
	End() error
	Accept(v reflect.Value) error

	Result() reflect.Value

	SetState(int)
	GetState() int
}

type BaseValve struct {
	fn    reflect.Value
	next  FuncValve
	state int
}

func CheckState(state, flag int) bool {
	return state&flag != 0
}

func SetState(state, flag int) int {
	return state | flag
}

func UnsetState(state, flag int) int {
	return state & (^flag)
}

func (valve *BaseValve) Init(fn interface{}) error {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		return errors.New("param is not a function")
	}
	valve.fn = v
	return nil
}

func (valve *BaseValve) Next(v FuncValve) {
	valve.next = v
}
func (valve *BaseValve) Reset() {
	if valve.next != nil {
		valve.next.Reset()
	}
}
func (valve *BaseValve) SetState(state int) {
	valve.state = state
}
func (valve *BaseValve) GetState() int {
	return valve.state
}
func (valve *BaseValve) Verify(t reflect.Type) bool {
	panic("cannot be here")
}
func (valve *BaseValve) Begin() {
	panic("cannot be here")
}
func (valve *BaseValve) End() {
	panic("cannot be here")
}
func (valve *BaseValve) Accept(v reflect.Value) {
	panic("cannot be here")
}
