// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/stream"
	"reflect"
	"strconv"
	"testing"
)

func TestOptionIsNil(t *testing.T) {
	o := stream.CanNil(nil)
	if o.IsPresent() {
		t.Fatal("option is nil")
	}
	if !o.IsNil() {
		t.Fatal("option is nil")
	}
}

func TestOptionIsPresent(t *testing.T) {
	o := stream.Some("some")
	if !o.IsPresent() {
		t.Fatal("option is nil")
	}
	if o.IsNil() {
		t.Fatal("option is nil")
	}
}

func TestOptionBind(t *testing.T) {
	opt := stream.Some(1)
	expect := stream.Some(2)
	double := func(in int) int {
		return in * 2
	}
	ret := opt.Bind(double)
	if !reflect.DeepEqual(expect, ret) {
		t.Fatalf("Option.Bind() failed: expected %v got %v", expect, ret)
	}
}

func TestOptionAnd(t *testing.T) {
	opt := stream.None
	expect := stream.None
	opt2 := stream.Some(1)
	ret := opt.And(opt2)
	if !reflect.DeepEqual(ret, expect) {
		t.Fatalf("Option.And() failed: expected %v got %v", expect, ret)
	}
}

func TestOptionAndThen(t *testing.T) {
	opt1 := stream.Some(1)
	double := func(in int) int {
		return in * 2
	}
	ret1 := opt1.AndThen(double)
	expect1 := stream.Some(2)
	if !reflect.DeepEqual(ret1, expect1) {
		t.Fatalf("Option.AndThen() failed: expected %v got %v", expect1, ret1)
	}
	opt2 := stream.None
	ret2 := opt2.AndThen(double)
	expect2 := stream.None
	if !reflect.DeepEqual(ret2, expect2) {
		t.Fatalf("Option.AndThen() failed: expected %v got %v", expect2, ret2)
	}
}

func TestOptionGet(t *testing.T) {
	opt1 := stream.Some(1)
	expect := 1
	if !reflect.DeepEqual(expect, opt1.Get()) {
		t.Fatalf("Option.Get() failed: expected %v got %v", expect, opt1.Get())
	}
}

func TestOptionElse(t *testing.T) {
	opt := stream.None
	expect := stream.Some(1)
	opt2 := stream.Some(1)
	ret := opt.Else(opt2)
	if !reflect.DeepEqual(ret, expect) {
		t.Fatalf("Option.Else() failed: expected %v got %v", expect, ret)
	}
}

func TestOptionFilter(t *testing.T) {
	opt := stream.Some(1)
	expect := stream.Some(1)
	ret := opt.Filter(func(i int) bool {
		if i == 1 {
			return true
		}
		return false
	})
	if !reflect.DeepEqual(ret, expect) {
		t.Fatalf("Option.Filter() failed: expected %v got %v", expect, ret)
	}
}

func TestOptionMap(t *testing.T) {
	opt := stream.Some(1)
	expect := stream.Some("1")
	ret := opt.Map(func(i int) interface{} {
		return strconv.Itoa(i)
	})
	if !reflect.DeepEqual(ret, expect) {
		t.Fatalf("Option.Map() failed: expected %v got %v", expect, ret)
	}
}

func TestOptionCall(t *testing.T) {
	opt := stream.Some(1)
	v := opt.Else(stream.Some(2)).Get()
	if v.(int) != 1 {
		t.Fatalf("expected %d got %d ", 2, v.(int))
	}

	v = opt.Map(func(i int) interface{} {
		return nil
	}).Else(stream.Some(2)).Get()
	if v.(int) != 2 {
		t.Fatalf("expected %d got %d ", 2, v.(int))
	}
}
