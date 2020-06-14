// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package stream

import (
	"github.com/xfali/stream/funcutil"
	"reflect"
)

type Collection interface {
	ElemType() reflect.Type
	Size() int
	Iterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() reflect.Value
}

type slice struct {
	slice reflect.Value
}

func NewSlice(o interface{}) *slice {
	if !funcutil.VerifySlice(o) {
		panic("interface not a slice")
	}
	return &slice {
		slice:reflect.ValueOf(o),
	}
}

func (c *slice) ElemType() reflect.Type {
	return c.slice.Type().Elem()
}

func (c *slice) Size() int {
	return c.slice.Len()
}

func (c *slice) Iterator() Iterator {
	return &sliceIterator{
		slice: c.slice,
		cur:   0,
	}
}

type sliceIterator struct {
	slice reflect.Value
	cur   int
}

func (c *sliceIterator) HasNext() bool {
	if c.cur < c.slice.Len() {
		return true
	}
	return false
}

func (c *sliceIterator) Next() reflect.Value {
	v := c.slice.Index(c.cur)
	c.cur++
	return v
}
