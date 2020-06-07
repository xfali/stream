// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package stream

import (
	"errors"
	"github.com/choleraehyq/gofunctools/functools"
	"math/rand"
	"reflect"
)

type SliceStream struct {
	slice interface{}
}

func Slice(slice interface{}) *SliceStream {
	if !VerifySlice(slice) {
		return nil
	}
	return &SliceStream{
		slice: slice,
	}
}

func (s *SliceStream) Count() int {
	v := reflect.ValueOf(s.slice)
	return v.Len()
}

func (s *SliceStream) Filter(fn interface{}) Stream {
	ret, err := functools.Filter(fn, s.slice)
	if err != nil {
		panic(err)
	}
	return &SliceStream{
		slice: ret,
	}
}

func (s *SliceStream) Limit(size int) Stream {
	ret, err := Limit(size, s.slice)
	if err != nil {
		panic(err)
	}
	return &SliceStream{
		slice: ret,
	}
}

func (s *SliceStream) Skip(size int) Stream {
	ret, err := Skip(size, s.slice)
	if err != nil {
		panic(err)
	}
	return &SliceStream{
		slice: ret,
	}
}

func (s *SliceStream) FindFirst() *Option {
	v := reflect.ValueOf(s.slice)
	if v.IsNil() || v.Len() == 0 {
		return None
	}

	first := v.Index(0)
	return CanNil(first.Interface())
}

func (s *SliceStream) FindLast() *Option {
	v := reflect.ValueOf(s.slice)
	if v.IsNil() || v.Len() == 0 {
		return None
	}

	last := v.Index(v.Len() - 1)
	return CanNil(last.Interface())
}

func (s *SliceStream) FindAny() *Option {
	v := reflect.ValueOf(s.slice)
	if v.IsNil() || v.Len() == 0 {
		return None
	}

	index := rand.Intn(v.Len())
	ret := v.Index(index)

	return CanNil(ret.Interface())
}

func (s *SliceStream) Foreach(eachFn interface{}) {
	in := reflect.ValueOf(s.slice)
	fn := reflect.ValueOf(eachFn)
	inType := in.Type().Elem()
	if !verifyForeachFuncType(fn, inType) {
		panic(errors.New("foreach Function must be of type func(" + inType.String() + ")"))
	}
	var param [1]reflect.Value
	for i := 0; i < in.Len(); i++ {
		param[0] = in.Index(i)
		fn.Call(param[:])
	}
}

func (s *SliceStream) AnyMatch(fn interface{}) bool {
	ret, err := functools.Any(fn, s.slice)
	if err != nil {
		panic(err)
	}
	return ret
}

func (s *SliceStream) AllMatch(fn interface{}) bool {
	ret, err := functools.All(fn, s.slice)
	if err != nil {
		panic(err)
	}
	return ret
}

func (s *SliceStream) Map(fn interface{}) Stream {
	ret, err := functools.Apply(fn, s.slice)
	if err != nil {
		panic(err)
	}
	return &SliceStream{
		slice: ret,
	}
}

func (s *SliceStream) Reduce(fn interface{}) interface{} {
	ret, err := Reduce(fn, s.slice)
	if err != nil {
		panic(err)
	}
	return ret
}

func (s *SliceStream) Collect() interface{} {
	return s.slice
}

func VerifySlice(o interface{}) bool {
	t := reflect.TypeOf(o)
	if t.Kind() != reflect.Slice {
		return false
	}
	return true
}

func verifyForeachFuncType(fn reflect.Value, elemType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 0 {
		return false
	}
	if fn.Type().In(0) != elemType {
		return false
	}
	return true
}
