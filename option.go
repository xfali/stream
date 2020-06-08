// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package stream

import "reflect"

type Option struct {
	val interface{}
}

var None = &Option{nil}

func Some(i interface{}) *Option {
	if i == nil {
		panic("Some() cannot get nil as argument. Please use None.")
	}
	return &Option{i}
}

func CanNil(i interface{}) *Option {
	return &Option{i}
}

func (o *Option) IsPresent() bool {
	if o.val != nil {
		return true
	}
	return false
}

func (o *Option) IsNil() bool {
	if o.val == nil {
		return true
	}
	return false
}

func (o *Option) IfPresent(fn func(interface{})) {
	if o.val != nil {
		fn(o.val)
	}
}

// If the Option is None, it will panic.
func (o *Option) Get() interface{} {
	if !o.IsPresent() {
		panic("Option: option is none")
	}
	return o.val
}

func (o *Option) Bind(function interface{}) *Option {
	fn := reflect.ValueOf(function)
	if !o.verifyBindFuncType(fn) {
		panic("Bind: Function must be of type func (valType) Option")
	}
	if !o.IsPresent() {
		return None
	}
	var param [1]reflect.Value
	param[0] = reflect.ValueOf(o.Get())
	out := fn.Call(param[:])[0]
	return &Option{out.Interface()}
}

// And returns None if this option is None, otherwise returns the option received.
func (o *Option) And(other *Option) *Option {
	if !o.IsPresent() {
		return None
	}
	return other
}

// And returns None if this option is None, otherwise returns the option received.
func (o *Option) Else(other *Option) *Option {
	if o.IsPresent() {
		return o
	}
	return other
}

// AndThen returns None if this option is None, otherwise calls the received function with the wrapped value and returns the result Option.
func (o *Option) AndThen(function interface{}) *Option {
	if !o.IsPresent() {
		return None
	}
	return o.Bind(function)
}

//fn func(TYPE) bool
func (o *Option) Filter(fn interface{}) *Option {
	fnValue := reflect.ValueOf(fn)
	if !o.verifyFilterFuncType(fnValue) {
		panic("Filter: Function must be of type func (valType) bool")
	}
	if !o.IsPresent() {
		return None
	}
	var param [1]reflect.Value
	param[0] = reflect.ValueOf(o.val)
	if fnValue.Call(param[:])[0].Bool() {
		return o
	} else {
		return None
	}
}

//fn func(OLD_TYPE) NEW_TYPE
func (o *Option) Map(mapFunc interface{}) *Option {
	if !o.IsPresent() {
		return None
	}

	fn := reflect.ValueOf(mapFunc)
	if !o.verifyMapFuncType(fn) {
		panic("Filter: Function must be of type func (valType) newType")
	}

	var param [1]reflect.Value
	param[0] = reflect.ValueOf(o.val)
	out := fn.Call(param[:])[0]
	return CanNil(out.Interface())
}

func (o *Option) verifyBindFuncType(fn reflect.Value) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
		return false
	}
	val := reflect.ValueOf(o.val)
	if fn.Type().In(0) != val.Type() {
		return false
	}
	return true
}

func (o *Option) verifyFilterFuncType(fn reflect.Value) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
		return false
	}
	val := reflect.TypeOf(o.val)
	if fn.Type().In(0) != val || fn.Type().Out(0).Kind() != reflect.Bool {
		return false
	}
	return true
}

func (o *Option) verifyMapFuncType(fn reflect.Value) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
		return false
	}
	if fn.Type().In(0) != reflect.TypeOf(o.val) {
		return false
	}
	return true
}
