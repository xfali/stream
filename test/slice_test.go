// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package test

import (
	"github.com/xfali/stream"
	"reflect"
	"testing"
)

func TestSliceFirst(t *testing.T) {
	s := stream.Slice([]int{1, 2, 3, 4, 5})
	o := s.FindFirst()
	if !reflect.DeepEqual(1, o.Get()) {
		t.Fatalf("Stream.First() failed: expected %v got %v", 1, o.Get())
	}
}

func TestSliceFindAny(t *testing.T) {
	s := stream.Slice([]int{1, 2, 3, 4, 5})
	for i := 0; i < 10; i++ {
		t.Log(s.FindAny().Get())
	}
}

func TestSliceForeach(t *testing.T) {
	s := stream.Slice([]int{1, 2, 3, 4, 5})
	s.Foreach(func(i int) {
		t.Log(i)
	})
}

func TestSliceFilter(t *testing.T) {
	s := stream.Slice([]int{1, 2, 3, 4, 5})
	s.Filter(func(i int) bool {
		if i == 2 {
			return false
		}
		return true
	}).Foreach(func(i int) {
		if i == 2 {
			t.Fatal("filter 2 but got it")
		}
	})
}

func TestSliceLimit(t *testing.T) {
	s := stream.Slice([]int{1, 2, 3, 4, 5})
	s.Limit(2).Foreach(func(i int) {
		t.Log(i)
		if i == 3 {
			t.Fatal("Limit cannot be 3 but got it")
		}
	})
	s.Limit(0).Foreach(func(i int) {
		switch i {
		case 1, 2, 3, 4, 5:
			t.Fatal("cannot be here")
		default:
			t.Log(i)
		}
	})

	s.Limit(5).Foreach(func(i int) {
		switch i {
		case 1, 2, 3, 4, 5:
			t.Log(i)
		default:
			t.Fatal("cannot be here")
		}
	})
}

func TestSliceSkip(t *testing.T) {
	s := stream.Slice([]int{1, 2, 3, 4, 5})
	s.Skip(2).Foreach(func(i int) {
		t.Log(i)
		if i == 1 || i == 2 {
			t.Fatal("Skip 1、2 but got it")
		}
	})
	s.Skip(0).Foreach(func(i int) {
		switch i {
		case 1, 2, 3, 4, 5:
			t.Log(i)
		default:
			t.Fatal("cannot be here")
		}
	})

	s.Skip(5).Foreach(func(i int) {
		switch i {
		case 1, 2, 3, 4, 5:
			t.Fatal("cannot be here")
		default:
			t.Log(i)
		}
	})
}

func TestSliceLimitSkip(t *testing.T) {
	s := stream.Slice([]int{1, 2, 3, 4, 5})
	s.Skip(2).Limit(1).Foreach(func(i int) {
		t.Log(i)
		if i != 3 {
			t.Fatal("Skip 1、2 but got it")
		}
	})
}
