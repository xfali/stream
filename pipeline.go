// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package stream

import (
	"github.com/xfali/stream/valve"
	"reflect"
)

type PipeStream struct {
	slice interface{}
	head  valve.FuncValve
	v     valve.FuncValve
}

func Pipeline(slice interface{}) *PipeStream {
	if !VerifySlice(slice) {
		return nil
	}
	n := &valve.NoneValve{}
	n.SetState(valve.NORMAL)
	return &PipeStream{
		slice: slice,
		head:  n,
		v:     n,
	}
}

func (s *PipeStream) Count() int {
	valve := &valve.CountValve{}
	s.v.Next(valve)
	s.v = valve
	return s.each().(int)
}

func (s *PipeStream) Limit(size int) Stream {
	valve := &valve.LimitValve{
		Limit: size,
	}
	s.v.Next(valve)
	s.v = valve
	return s
}

func (s *PipeStream) Filter(fn interface{}) Stream {
	valve := &valve.FilterValve{}
	valve.Init(fn)
	s.v.Next(valve)
	s.v = valve
	return s
}

func (s *PipeStream) Skip(size int) Stream {
	valve := &valve.SkipValve{
		Skip: size,
	}
	s.v.Next(valve)
	s.v = valve
	return s
}

func (s *PipeStream) Distinct(fn interface{}) Stream {
	valve := &valve.DistinctValve{}
	valve.Init(fn)
	s.v.Next(valve)
	s.v = valve
	return s
}

func (s *PipeStream) Sort(fn interface{}) Stream {
	valve := &valve.SortValve{}
	valve.Init(fn)
	s.v.Next(valve)
	s.v = valve
	return s
}

func (s *PipeStream) FindFirst() *Option {
	valve := &valve.FindFirstValve{}
	s.v.Next(valve)
	s.v = valve

	v := s.each()
	return CanNil(v)
}

func (s *PipeStream) FindLast() *Option {
	valve := &valve.FindLastValve{}
	s.v.Next(valve)
	s.v = valve

	v := s.each()
	return CanNil(v)
}

func (s *PipeStream) FindAny() *Option {
	valve := &valve.FindAnyValve{}
	s.v.Next(valve)
	s.v = valve

	v := s.each()
	return CanNil(v)
}

func (s *PipeStream) Foreach(eachFn interface{}) {
	valve := &valve.ForeachValve{}
	valve.Init(eachFn)
	s.v.Next(valve)
	s.v = valve

	s.each()
}

func (s *PipeStream) Peek(eachFn interface{}) Stream {
	valve := &valve.ForeachValve{}
	valve.Init(eachFn)
	s.v.Next(valve)
	s.v = valve

	s.each()

	return s
}

func (s *PipeStream) AnyMatch(fn interface{}) bool {
	valve := &valve.MatchAnyValve{}
	valve.Init(fn)
	s.v.Next(valve)
	s.v = valve

	return s.each().(bool)
}

func (s *PipeStream) AllMatch(fn interface{}) bool {
	valve := &valve.MatchAllValve{}
	valve.Init(fn)
	s.v.Next(valve)
	s.v = valve

	return s.each().(bool)
}

func (s *PipeStream) Map(fn interface{}) Stream {
	valve := &valve.MapValve{}
	valve.Init(fn)
	s.v.Next(valve)
	s.v = valve

	return s
}

func (s *PipeStream) FlatMap(fn interface{}) Stream {
	valve := &valve.FlatMapValve{}
	valve.Init(fn)
	s.v.Next(valve)
	s.v = valve

	return s
}

func (s *PipeStream) Reduce(fn, initValue interface{}) interface{} {
	valve := &valve.ReduceValve{}
	if initValue != nil {
		valve.V = reflect.ValueOf(initValue)
	}
	valve.Init(fn)
	s.v.Next(valve)
	s.v = valve

	return s.each()
}

func (s *PipeStream) Collect() interface{} {
	valve := &valve.CollectValve{}
	s.v.Next(valve)
	s.v = valve

	return s.each()
}

func (s *PipeStream) each() interface{} {
	in := reflect.ValueOf(s.slice)
	inType := in.Type().Elem()
	err := s.head.Verify(inType)
	if err != nil {
		panic(err)
	}
	s.head.Begin(in.Len())
	for i := 0; i < in.Len(); i++ {
		s.head.Accept(in.Index(i))
	}
	s.head.End()
	v := s.head.Result()
	if v.IsValid() {
		return v.Interface()
	} else {
		return nil
	}
}
