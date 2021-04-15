// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	var i interface{} = 1
	inType := reflect.TypeOf(i)
	t.Log(inType)
	v := reflect.ValueOf(i).Convert(reflect.TypeOf(1))
	t.Log(v)
}

func test(o ...interface{}) reflect.Value {
	t := reflect.SliceOf(reflect.TypeOf(o[0]))
	sl := reflect.MakeSlice(t, 0, len(o))
	for _, v := range o {
		sl = reflect.Append(sl, reflect.ValueOf(v))
	}
	return sl
}

func TestReflect2(t *testing.T) {
	sl := test(1, 2, 3, 4)
	for i := 0; i < sl.Len(); i++ {
		t.Log(sl.Index(i).Interface())
	}
}

func TestReflectValid(t *testing.T) {
	v := reflect.Value{}
	t.Log(v.IsValid())
	if v.IsValid() {
		t.Fatal("cannot be valid")
	}
}
