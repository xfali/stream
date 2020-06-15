// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package valve

import (
	"errors"
	"github.com/xfali/stream/collector"
	"reflect"
)

type CollectValve struct {
	BaseValve
	collector collector.Collector
	elemType  reflect.Type
}

func (valve *CollectValve) Verify(t reflect.Type) error {
	valve.elemType = t
	return nil
}

func (valve *CollectValve) Init(fn interface{}) error {
	if fn != nil {
		c, ok := fn.(collector.Collector)
		if !ok {
			return errors.New("param is not collector.Collector")
		}
		valve.collector = c
	} else {
		valve.collector = collector.ToSlice()
	}
	return nil
}

func (valve *CollectValve) Begin(count int) error {
	if count == -1 {
		count = DefaultCapacity
	}
	valve.collector.New(valve.elemType, count)
	return nil
}

func (valve *CollectValve) End() error {
	return nil
}

func (valve *CollectValve) Accept(v reflect.Value) error {
	return valve.collector.Add(v)
}

func (valve *CollectValve) Result() reflect.Value {
	return valve.collector.Result()
}
