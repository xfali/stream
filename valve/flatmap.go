// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package valve

import (
	"errors"
	"github.com/xfali/stream/funcutil"
	"reflect"
)

type FlatMapValve struct {
	BaseValve
}

func (valve *FlatMapValve) Verify(t reflect.Type) error {
	if !funcutil.VerifyFlatMapFunction(valve.fn, t) {
		return errors.New("flatmap: Function must be of type func(" + t.String() + ") []interface{}")
	}
	return valve.next.Verify(valve.fn.Type().Out(0).Elem())
}

func (valve *FlatMapValve) Begin(count int) error {
	valve.next.SetState(valve.state)
	return valve.next.Begin(-1)
}

func (valve *FlatMapValve) End() error {
	return nil
}

func (valve *FlatMapValve) Accept(v reflect.Value) error {
	var param [1]reflect.Value
	param[0] = v
	newSlice := valve.fn.Call(param[:])[0]
	for j := 0; j < newSlice.Len(); j++ {
		err := valve.next.Accept(newSlice.Index(j))
		if err != nil {
			return err
		}
	}
	return nil
}

func (valve *FlatMapValve) Result() reflect.Value {
	return valve.next.Result()
}
