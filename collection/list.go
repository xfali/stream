// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package collection

import (
	list2 "container/list"
	"reflect"
)

type list struct {
	l *list2.List
}

func NewList(l *list2.List) *list {
	return &list{
		l: l,
	}
}

func CreateList(v ...interface{}) *list {
	if len(v) == 0 {
		return nil
	}

	l := list2.New()
	for i := range v {
		l.PushBack(v[i])
	}
	return &list{
		l: l,
	}
}

func (c *list) ElemType() reflect.Type {
	o := c.l.Front()
	if o == nil {
		return nil
	}
	return reflect.TypeOf(o.Value)
}

func (c *list) Size() int {
	return c.l.Len()
}

func (c *list) Iterator() Iterator {
	return &listIterator{
		e:   c.l.Front(),
		cur: 0,
	}
}

func (c *list) Add(value reflect.Value) {
	c.l.PushBack(value.Interface())
}

type listIterator struct {
	e   *list2.Element
	cur int
}

func (c *listIterator) HasNext() bool {
	return c.e != nil
}

func (c *listIterator) Next() reflect.Value {
	v := reflect.ValueOf(c.e.Value)
	c.e = c.e.Next()
	return v
}
