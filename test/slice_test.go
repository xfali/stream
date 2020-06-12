// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package test

import (
	"github.com/xfali/stream"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestSliceFirst(t *testing.T) {
	s := stream.Slice([]int{1, 2, 3, 4, 5})
	o := s.FindFirst()
	if !reflect.DeepEqual(1, o.Get()) {
		t.Fatalf("Stream.First() failed: expected %v got %v", 1, o.Get())
	}
}

func TestSliceLast(t *testing.T) {
	s := stream.Slice([]int{1, 2, 3, 4, 5})
	o := s.FindLast()
	if !reflect.DeepEqual(5, o.Get()) {
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

func TestSliceDistinct(t *testing.T) {
	s := stream.Slice([]int{1, 2, 2, 4, 5})
	x := s.Distinct(func(a, b int) bool {
		return a == b
	}).Collect().([]int)
	if x[2] == 2 {
		t.Fatal("cannot be 2")
	}

	s.Distinct(func(a, b int) bool {
		return a == b
	}).Foreach(func(i int) {
		t.Log(i)
	})
}

func TestSliceSort(t *testing.T) {
	t.Run("asc", func(t *testing.T) {
		s := stream.Slice([]int{5, 2, 3, 1, 4})
		s.Sort(func(a, b int) int {
			return a - b
		}).Foreach(func(i int) {
			t.Log(i)
		})
	})

	t.Run("desc", func(t *testing.T) {
		s := stream.Slice([]int{5, 2, 3, 1, 4})
		s.Sort(func(a, b int) int {
			return b - a
		}).Foreach(func(i int) {
			t.Log(i)
		})
	})

	t.Run("asc repeat", func(t *testing.T) {
		s := stream.Slice([]int{5, 2, 5, 2, 3, 1, 4})
		s.Sort(func(a, b int) int {
			return a - b
		}).Foreach(func(i int) {
			t.Log(i)
		})
	})
}

func TestSliceMap(t *testing.T) {
	s := stream.Slice([]string{"1", "2", "3"})
	s.Map(func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}).Foreach(func(i int) {
		t.Log(i)
	})
}

func TestSliceFlatMap(t *testing.T) {
	s := stream.Slice([]string{"hello world", "xfali stream"})
	s.FlatMap(func(s string) []string {
		return strings.Split(s, " ")
	}).Foreach(func(i string) {
		t.Log(i)
	})

	s = stream.Slice([]string{"1,2,3,4", "5,6,7,8"})
	s.FlatMap(func(s string) []int {
		return stream.Slice(strings.Split(s, ",")).Map(func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		}).Collect().([]int)
	}).Foreach(func(i int) {
		t.Log(i)
	})
}

func TestSliceReduce(t *testing.T) {
	t.Run("without init value", func(t *testing.T) {
		s := stream.Slice([]int{1, 2, 3, 4, 5})
		ret := s.Reduce(func(d, o int) int {
			return d + o
		}, nil).(int)
		if ret != 15 {
			t.Fatal("expect 15 but get: ", ret)
		}
	})

	t.Run("with init value", func(t *testing.T) {
		s := stream.Slice([]int{1, 2, 3, 4, 5})
		ret := s.Reduce(func(d, o int) int {
			return d + o
		}, 2).(int)
		if ret != 17 {
			t.Fatal("expect 15 but get: ", ret)
		}
	})

	t.Run("with string init value", func(t *testing.T) {
		s := stream.Slice([]string{"w", "o", "r", "l", "d"})
		ret := s.Reduce(func(d, o string) string {
			return d + o
		}, "hello ").(string)
		if ret != "hello world" {
			t.Fatal("expect 15 but get: ", ret)
		}
	})
}

type student struct {
	name   string
	gender string
}

func TestSliceGroup(t *testing.T) {
	s := []student{
		{"lilei", "boy"},
		{"hanmeimei", "girl"},
		{"lily", "girl"},
		{"jim", "boy"},
		{"tom", "boy"},
	}

	ret, err := stream.Group(func(s student) (string, student) {
		return s.gender, s
	}, s)
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range ret.(map[string][]student) {
		for _, stu := range v {
			t.Log(k, stu)
		}
	}
}
