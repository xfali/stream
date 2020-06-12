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

type ReduceValve struct {
	BaseValve
	V reflect.Value
}

func (valve *ReduceValve) Reset() {
	valve.V = reflect.Value{}
}

func (valve *ReduceValve) Verify(t reflect.Type) error {
	if ! funcutil.VerifyReduceFuncType(valve.fn, t) {
		return errors.New("reduce: Function must be of type func(" + t.String() + ")" + t.String())
	}
	return nil
}

func (valve *ReduceValve) Begin(count int) error {
	return nil
}

func (valve *ReduceValve) End() error {
	return nil
}

func (valve *ReduceValve) Accept(v reflect.Value) error {
	if !valve.V.IsValid() {
		valve.V = v
	} else {
		var param [2]reflect.Value
		param[0] = valve.V
		param[1] = v
		valve.V = valve.fn.Call(param[:])[0]
	}
	return nil
}

func (valve *ReduceValve) Result() reflect.Value {
	return valve.V
}
