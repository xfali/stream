// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package stream

type Stream interface {
    Count() int

    FindFirst() *Option

    FindAny() *Option

    //fn func(TYPE) bool
    Filter(fn interface{}) Stream

    //fn func(OLD_TYPE) NEW_TYPE
    Map(fn interface{}) Stream

    //fn func(TYPE)
    Foreach(fn interface{})

    //fn func(TYPE) bool
    AnyMatch(fn interface{}) bool

    //fn func(o TYPE) bool
    AllMatch(fn interface{}) bool

    //fn func(out, in TYPE) interface{}
    Reduce(fn interface{}) interface{}

    Collect() interface{}
}
