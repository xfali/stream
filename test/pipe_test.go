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

func TestPipeFilter(t *testing.T) {
	s := stream.Pipeline([]int{1, 2, 3, 4, 5})
	s.Filter(func(i int) bool {
		if i == 2 {
			return false
		}
		return true
	}).Foreach(func(i int) {
		if i == 2 {
			t.Fatal("filter 2 but got it")
		} else {
			t.Log(i)
		}
	})
}

func TestPipeFirst(t *testing.T) {
	s := stream.Pipeline([]int{1, 2, 3, 4, 5})
	o := s.FindFirst()
	if !reflect.DeepEqual(1, o.Get()) {
		t.Fatalf("Stream.First() failed: expected %v got %v", 1, o.Get())
	}
}

func TestPipeLast(t *testing.T) {
	s := stream.Pipeline([]int{1, 2, 3, 4, 5})
	o := s.FindLast()
	if !reflect.DeepEqual(5, o.Get()) {
		t.Fatalf("Stream.First() failed: expected %v got %v", 1, o.Get())
	}
}

func TestPipeFindAny(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := stream.Pipeline([]int{1, 2, 3, 4, 5})
		t.Log(s.FindAny().Get())
	}
	for i := 0; i < 10; i++ {
		s := stream.Pipeline([]int{1, 2, 3, 4, 5})
		v := s.Filter(func(i int) bool {
			return i != 2
		}).FindAny().Get()
		t.Log(v)
		if v == 2 {
			t.Fatal("cannot be 2")
		}
	}
}

func TestPipeLimit(t *testing.T) {
	s := stream.Pipeline([]int{1, 2, 3, 4, 5})
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

func TestPipeSkip(t *testing.T) {
	t.Run("skip 2", func(t *testing.T) {
		stream.Pipeline([]int{1, 2, 3, 4, 5}).Skip(2).Foreach(func(i int) {
			t.Log(i)
			if i == 1 || i == 2 {
				t.Fatal("Skip 1、2 but got it")
			}
		})
	})

	t.Run("skip 0", func(t *testing.T) {
		stream.Pipeline([]int{1, 2, 3, 4, 5}).Skip(0).Foreach(func(i int) {
			switch i {
			case 1, 2, 3, 4, 5:
				t.Log(i)
			default:
				t.Fatal("cannot be here")
			}
		})
	})

	t.Run("skip 5", func(t *testing.T) {
		stream.Pipeline([]int{1, 2, 3, 4, 5}).Skip(5).Foreach(func(i int) {
			switch i {
			case 1, 2, 3, 4, 5:
				t.Fatal("cannot be here")
			default:
				t.Log(i)
			}
		})
	})
}

func TestPipeLimitSkip(t *testing.T) {
	s := stream.Pipeline([]int{1, 2, 3, 4, 5})
	s.Skip(2).Limit(1).Foreach(func(i int) {
		t.Log(i)
		if i != 3 {
			t.Fatal("Skip 1、2 but got it")
		}
	})
}

func TestPipeDistinct(t *testing.T) {
	s := stream.Pipeline([]int{1, 2, 2, 4, 5})
	s.Distinct(func(a, b int) int {
		return a - b
	}).Foreach(func(i int) {
		t.Log(i)
	})
}

func TestPipeSort(t *testing.T) {
	t.Run("asc", func(t *testing.T) {
		s := stream.Pipeline([]int{5, 2, 3, 1, 4})
		s.Sort(func(a, b int) int {
			return a - b
		}).Foreach(func(i int) {
			t.Log(i)
		})
	})

	t.Run("desc", func(t *testing.T) {
		s := stream.Pipeline([]int{5, 2, 3, 1, 4})
		s.Sort(func(a, b int) int {
			return b - a
		}).Foreach(func(i int) {
			t.Log(i)
		})
	})

	t.Run("asc repeat", func(t *testing.T) {
		s := stream.Pipeline([]int{5, 2, 5, 2, 3, 1, 4})
		s.Sort(func(a, b int) int {
			return a - b
		}).Foreach(func(i int) {
			t.Log(i)
		})
	})
}

func TestPipeMap(t *testing.T) {
	stream.Pipeline([]string{"1", "2", "3"}).Map(func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}).Foreach(func(i int) {
		switch i {
		case 1, 2, 3:
			t.Log(i)
		default:
			t.Fatal("cannot be here")
		}
	})
}

func TestPipeFlatMap(t *testing.T) {
	stream.Pipeline([]string{"hello world", "xfali stream"}).FlatMap(func(s string) []string {
		return strings.Split(s, " ")
	}).Foreach(func(i string) {
		switch i {
		case "hello", "world", "xfali", "stream":
			t.Log(i)
		default:
			t.Fatal("cannot be here")
		}
	})

	stream.Pipeline([]string{"1,2,3,4", "5,6,7,8"}).FlatMap(func(s string) []int {
		return stream.Pipeline(strings.Split(s, ",")).Map(func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		}).Collect().([]int)
	}).Foreach(func(i int) {
		switch i {
		case 1, 2, 3, 4, 5, 6, 7, 8:
			t.Log(i)
		default:
			t.Fatal("cannot be here")
		}
	})
}