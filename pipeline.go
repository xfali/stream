// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package stream

import (
	list2 "container/list"
	"github.com/xfali/stream/collection"
	"github.com/xfali/stream/valve"
	"reflect"
)

type PipelineStream struct {
	c    collection.Collection
	head valve.FuncValve
	v    valve.FuncValve
}

func Slice(slice interface{}) *PipelineStream {
	return New(collection.NewSlice(slice))
}

func List(list *list2.List) *PipelineStream {
	return New(collection.NewList(list))
}

func New(c collection.Collection) *PipelineStream {
	n := &valve.NoneValve{}
	n.SetState(valve.NORMAL)
	return &PipelineStream{
		c:    c,
		head: n,
		v:    n,
	}
}

func (s *PipelineStream) Count() int {
	valve := &valve.CountValve{}
	s.v.Next(valve)
	s.v = valve
	return s.each().(int)
}

func (s *PipelineStream) Limit(size int) Stream {
	valve := &valve.LimitValve{
		Limit: size,
	}
	s.v.Next(valve)
	s.v = valve
	return s
}

func (s *PipelineStream) Filter(fn interface{}) Stream {
	valve := &valve.FilterValve{}
	err := valve.Init(fn)
	if err != nil {
		panic(err)
	}
	s.v.Next(valve)
	s.v = valve
	return s
}

func (s *PipelineStream) Skip(size int) Stream {
	valve := &valve.SkipValve{
		Skip: size,
	}
	s.v.Next(valve)
	s.v = valve
	return s
}

func (s *PipelineStream) Distinct(fn interface{}) Stream {
	valve := &valve.DistinctValve{}
	err := valve.Init(fn)
	if err != nil {
		panic(err)
	}
	s.v.Next(valve)
	s.v = valve
	return s
}

func (s *PipelineStream) Sort(fn interface{}) Stream {
	valve := &valve.SortValve{}
	err := valve.Init(fn)
	if err != nil {
		panic(err)
	}
	s.v.Next(valve)
	s.v = valve
	return s
}

func (s *PipelineStream) FindFirst() *Option {
	valve := &valve.FindFirstValve{}
	s.v.Next(valve)
	s.v = valve

	v := s.each()
	return CanNil(v)
}

func (s *PipelineStream) FindLast() *Option {
	valve := &valve.FindLastValve{}
	s.v.Next(valve)
	s.v = valve

	v := s.each()
	return CanNil(v)
}

func (s *PipelineStream) FindAny() *Option {
	valve := &valve.FindAnyValve{}
	s.v.Next(valve)
	s.v = valve

	v := s.each()
	return CanNil(v)
}

func (s *PipelineStream) Foreach(eachFn interface{}) {
	valve := &valve.ForeachValve{}
	err := valve.Init(eachFn)
	if err != nil {
		panic(err)
	}
	s.v.Next(valve)
	s.v = valve

	s.each()
}

func (s *PipelineStream) Peek(eachFn interface{}) Stream {
	valve := &valve.PeekValve{}
	err := valve.Init(eachFn)
	if err != nil {
		panic(err)
	}
	s.v.Next(valve)
	s.v = valve
	return s
}

func (s *PipelineStream) AnyMatch(fn interface{}) bool {
	valve := &valve.MatchAnyValve{}
	err := valve.Init(fn)
	if err != nil {
		panic(err)
	}
	s.v.Next(valve)
	s.v = valve

	return s.each().(bool)
}

func (s *PipelineStream) AllMatch(fn interface{}) bool {
	valve := &valve.MatchAllValve{}
	err := valve.Init(fn)
	if err != nil {
		panic(err)
	}
	s.v.Next(valve)
	s.v = valve

	return s.each().(bool)
}

func (s *PipelineStream) Map(fn interface{}) Stream {
	valve := &valve.MapValve{}
	err := valve.Init(fn)
	if err != nil {
		panic(err)
	}
	s.v.Next(valve)
	s.v = valve

	return s
}

func (s *PipelineStream) FlatMap(fn interface{}) Stream {
	valve := &valve.FlatMapValve{}
	err := valve.Init(fn)
	if err != nil {
		panic(err)
	}
	s.v.Next(valve)
	s.v = valve

	return s
}

func (s *PipelineStream) Reduce(fn, initValue interface{}) interface{} {
	valve := &valve.ReduceValve{}
	if initValue != nil {
		valve.V = reflect.ValueOf(initValue)
	}
	err := valve.Init(fn)
	if err != nil {
		panic(err)
	}
	s.v.Next(valve)
	s.v = valve

	return s.each()
}

func (s *PipelineStream) Collect(collector interface{}) interface{} {
	valve := &valve.CollectValve{}
	err := valve.Init(collector)
	if err != nil {
		panic(err)
	}
	s.v.Next(valve)
	s.v = valve

	return s.each()
}

func (s *PipelineStream) each() interface{} {
	inType := s.c.ElemType()
	if inType == nil {
		return nil
	}
	err := s.head.Verify(inType)
	if err != nil {
		panic(err)
	}
	err = s.head.Begin(s.c.Size())
	if err != nil {
		panic(err)
	}
	iter := s.c.Iterator()
	for iter.HasNext() {
		v := iter.Next()
		err = s.head.Accept(v)
		if err != nil {
			panic(err)
		}
	}
	err = s.head.End()
	if err != nil {
		panic(err)
	}
	v := s.head.Result()
	if v.IsValid() {
		return v.Interface()
	} else {
		return nil
	}
}
