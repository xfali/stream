// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package test

import (
	"container/list"
	"github.com/xfali/stream"
	"github.com/xfali/stream/collection"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

var newFunc = NewSlice

func NewSlice(o ...interface{}) stream.Stream {
	return stream.New(collection.CreateSlice(o...))
}

func NewList(o ...interface{}) stream.Stream {
	return stream.New(collection.CreateList(o...))
}

func TestSlice(t *testing.T) {
	stream.Slice([]string{"0,1,3,2,2,4", "5,8,7,6,7,9"}).FlatMap(func(s string) []int {
		return stream.Slice(strings.Split(s, ",")).Map(func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		}).Collect().([]int)
	}).Filter(func(i int) bool {
		return i != 5
	}).Sort(func(a, b int) int {
		return a - b
	}).Distinct(func(a, b int) bool {
		return a == b
	}).Map(func(i int) string {
		return strconv.Itoa(i)
	}).Foreach(func(s string) {
		t.Log(s)
	})
}

func TestList(t *testing.T) {
	testList := list.New()
	testList.PushBack("0,1,3,2,2,4")
	testList.PushBack("5,8,7,6,7,9")
	stream.List(testList).FlatMap(func(s string) []int {
		return stream.Slice(strings.Split(s, ",")).Map(func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		}).Collect().([]int)
	}).Filter(func(i int) bool {
		return i != 5
	}).Sort(func(a, b int) int {
		return a - b
	}).Distinct(func(a, b int) bool {
		return a == b
	}).Map(func(i int) string {
		return strconv.Itoa(i)
	}).Foreach(func(s string) {
		t.Log(s)
	})
}

func TestStreamFilter(t *testing.T) {
	s := newFunc(1, 2, 3, 4, 5)
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

func TestStreamFirst(t *testing.T) {
	s := newFunc(1, 2, 3, 4, 5)
	o := s.FindFirst()
	if !reflect.DeepEqual(1, o.Get()) {
		t.Fatalf("Stream.First() failed: expected %v got %v", 1, o.Get())
	}
}

func TestStreamLast(t *testing.T) {
	s := newFunc(1, 2, 3, 4, 5)
	o := s.FindLast()
	if !reflect.DeepEqual(5, o.Get()) {
		t.Fatalf("Stream.First() failed: expected %v got %v", 1, o.Get())
	}
}

func TestStreamFindAny(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := newFunc(1, 2, 3, 4, 5)
		t.Log(s.FindAny().Get())
	}
	for i := 0; i < 10; i++ {
		s := newFunc(1, 2, 3, 4, 5)
		v := s.Filter(func(i int) bool {
			return i != 2
		}).FindAny().Get()
		t.Log(v)
		if v == 2 {
			t.Fatal("cannot be 2")
		}
	}
}

func TestStreamLimit(t *testing.T) {
	s := newFunc(1, 2, 3, 4, 5)
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

func TestStreamSkip(t *testing.T) {
	t.Run("skip 2", func(t *testing.T) {
		newFunc(1, 2, 3, 4, 5).Skip(2).Foreach(func(i int) {
			t.Log(i)
			if i == 1 || i == 2 {
				t.Fatal("Skip 1、2 but got it")
			}
		})
	})

	t.Run("skip 0", func(t *testing.T) {
		newFunc(1, 2, 3, 4, 5).Skip(0).Foreach(func(i int) {
			switch i {
			case 1, 2, 3, 4, 5:
				t.Log(i)
			default:
				t.Fatal("cannot be here")
			}
		})
	})

	t.Run("skip 5", func(t *testing.T) {
		newFunc(1, 2, 3, 4, 5).Skip(5).Foreach(func(i int) {
			switch i {
			case 1, 2, 3, 4, 5:
				t.Fatal("cannot be here")
			default:
				t.Log(i)
			}
		})
	})
}

func TestStreamLimitSkip(t *testing.T) {
	s := newFunc(1, 2, 3, 4, 5)
	s.Skip(2).Limit(1).Foreach(func(i int) {
		t.Log(i)
		if i != 3 {
			t.Fatal("Skip 1、2 but got it")
		}
	})
}

func TestStreamDistinct(t *testing.T) {
	t.Run("distinct", func(t *testing.T) {
		newFunc(1, 2, 2, 4, 5).Distinct(func(a, b int) bool {
			return a == b
		}).Foreach(func(i int) {
			t.Log(i)
		})
	})

	t.Run("distinct twice", func(t *testing.T) {
		newFunc(1, 2, 2, 4, 5).Distinct(func(a, b int) bool {
			return a == b
		}).Distinct(func(a, b int) bool {
			return a == b
		}).Foreach(func(i int) {
			t.Log(i)
		})
	})

	t.Run("distinct after sort", func(t *testing.T) {
		newFunc(1, 2, 2, 4, 5).Sort(func(a, b int) int {
			return a - b
		}).Distinct(func(a, b int) bool {
			return a == b
		}).Foreach(func(i int) {
			t.Log(i)
		})
	})
}

func TestStreamSort(t *testing.T) {
	t.Run("asc", func(t *testing.T) {
		var ret []int
		s := newFunc(5, 2, 3, 1, 4)
		s.Sort(func(a, b int) int {
			return a - b
		}).Foreach(func(i int) {
			ret = append(ret, i)
			t.Log(i)
		})
		for i := range ret {
			if ret[i] != i+1 {
				t.Fatal("sort error")
			}
		}
	})

	t.Run("desc", func(t *testing.T) {
		var ret []int
		s := newFunc(5, 2, 3, 1, 4)
		s.Sort(func(a, b int) int {
			return b - a
		}).Foreach(func(i int) {
			t.Log(i)
			ret = append(ret, i)
		})
		for i := range ret {
			if ret[len(ret)-i-1] != i+1 {
				t.Fatal("sort error")
			}
		}
	})

	t.Run("asc repeat", func(t *testing.T) {
		s := newFunc(5, 2, 5, 2, 3, 1, 4)
		s.Sort(func(a, b int) int {
			return a - b
		}).Foreach(func(i int) {
			t.Log(i)
		})
	})

	t.Run("desc twice", func(t *testing.T) {
		s := newFunc(5, 2, 3, 1, 4)
		s.Sort(func(a, b int) int {
			return b - a
		}).Sort(func(a, b int) int {
			return b - a
		}).Foreach(func(i int) {
			t.Log(i)
		})
	})
}

func TestStreamMap(t *testing.T) {
	newFunc("1", "2", "3").Map(func(s string) int {
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

func TestStreamMapStruct(t *testing.T) {
	t.Run("struct2struct", func(t *testing.T) {
		newFunc(student{
			name:   "lilei",
			gender: "boy",
			age:    10,
		}, student{
			name:   "hanmeimei",
			gender: "girl",
			age:    9,
		}, student{
			name:   "lucy",
			gender: "girl",
			age:    10,
		}, student{
			name:   "jim",
			gender: "boy",
			age:    10,
		}, student{
			name:   "lily",
			gender: "girl",
			age:    10,
		}).Map(func(s student) people {
			return people{
				name:       s.name,
				gender:     s.gender,
				profession: "student",
			}
		}).Foreach(func(i people) {
			switch i.name {
			case "lilei", "hanmeimei", "lucy", "lily", "jim":
				t.Log(i)
			default:
				t.Fatal()
			}
		})
	})

	t.Run("struct2int", func(t *testing.T) {
		newFunc(student{
			name:   "lilei",
			gender: "boy",
			age:    10,
		}, student{
			name:   "hanmeimei",
			gender: "girl",
			age:    9,
		}, student{
			name:   "lucy",
			gender: "girl",
			age:    10,
		}, student{
			name:   "jim",
			gender: "boy",
			age:    10,
		}, student{
			name:   "lily",
			gender: "girl",
			age:    10,
		}).Map(func(s student) int {
			return s.age
		}).Foreach(func(i int) {
			switch i {
			case 9, 10:
				t.Log(i)
			default:
				t.Fatal()
			}
		})
	})
}

func TestStreamFlatMap(t *testing.T) {
	newFunc("hello world", "xfali stream").FlatMap(func(s string) []string {
		return strings.Split(s, " ")
	}).Foreach(func(i string) {
		switch i {
		case "hello", "world", "xfali", "stream":
			t.Log(i)
		default:
			t.Fatal("cannot be here")
		}
	})

	newFunc("1,2,3,4", "5,6,7,8").FlatMap(func(s string) []int {
		return stream.Slice(strings.Split(s, ",")).Map(func(s string) int {
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

func TestStreamCount(t *testing.T) {
	t.Run("limit", func(t *testing.T) {
		c := newFunc(1, 2, 3, 4, 5).Limit(2).Count()
		if c != 2 {
			t.Fatal("expect 2 but get: ", c)
		}
	})

	t.Run("filter", func(t *testing.T) {
		c := newFunc(1, 2, 3, 4, 5).Filter(func(a int) bool {
			return a != 2
		}).Count()
		if c != 4 {
			t.Fatal("expect 2 but get: ", c)
		}
	})
}

func TestStreamForeach(t *testing.T) {
	t.Run("limit", func(t *testing.T) {
		newFunc(1, 2, 3, 4, 5).Limit(2).Foreach(func(i int) {
			switch i {
			case 1, 2:
				t.Log(i)
			default:
				t.Fatal("cannot be here")
			}
		})
	})

	t.Run("filter", func(t *testing.T) {
		newFunc(1, 2, 3, 4, 5).Filter(func(a int) bool {
			return a != 2
		}).Foreach(func(i int) {
			switch i {
			case 1, 3, 4, 5:
				t.Log(i)
			default:
				t.Fatal("cannot be here")
			}
		})
	})
}

func TestStreamPeek(t *testing.T) {
	t.Run("limit", func(t *testing.T) {
		c := newFunc(1, 2, 3, 4, 5).Limit(2).Peek(func(i int) {
			switch i {
			case 1, 2:
				t.Log(i)
			default:
				t.Fatal("cannot be here")
			}
		}).Count()
		if c != 2 {
			t.Fatal("expect 2 but get: ", c)
		}
	})

	t.Run("filter", func(t *testing.T) {
		c := newFunc(1, 2, 3, 4, 5).Filter(func(a int) bool {
			return a != 2
		}).Peek(func(i int) {
			switch i {
			case 1, 3, 4, 5:
				t.Log(i)
			default:
				t.Fatal("cannot be here")
			}
		}).Count()
		if c != 4 {
			t.Fatal("expect 2 but get: ", c)
		}
	})
}

func TestStreamAnyMatch(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		c := newFunc(1, 2, 3, 4, 5).AnyMatch(func(i int) bool {
			return i == 3
		})
		if !c {
			t.Fatal("expect true but get: ", c)
		}
	})

	t.Run("false", func(t *testing.T) {
		c := newFunc(1, 2, 3, 4, 5).AnyMatch(func(i int) bool {
			return i == 6
		})
		if c {
			t.Fatal("expect false but get: ", c)
		}
	})
}

func TestStreamAllMatch(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		c := newFunc(1, 2, 3, 4, 5).AllMatch(func(i int) bool {
			return i != 6
		})
		if !c {
			t.Fatal("expect true but get: ", c)
		}
	})

	t.Run("false", func(t *testing.T) {
		c := newFunc(1, 2, 3, 4, 5).AllMatch(func(i int) bool {
			return i != 4
		})
		if c {
			t.Fatal("expect false but get: ", c)
		}
	})
}

func TestStreamReduce(t *testing.T) {
	t.Run("without init value", func(t *testing.T) {
		ret := newFunc(1, 2, 3, 4, 5).Reduce(func(d, o int) int {
			return d + o
		}, nil).(int)
		if ret != 15 {
			t.Fatal("expect 15 but get: ", ret)
		}
	})

	t.Run("with init value", func(t *testing.T) {
		ret := newFunc(1, 2, 3, 4, 5).Reduce(func(d, o int) int {
			return d + o
		}, 2).(int)
		if ret != 17 {
			t.Fatal("expect 17 but get: ", ret)
		}
	})

	t.Run("with string init value", func(t *testing.T) {
		ret := newFunc("w", "o", "r", "l", "d").Reduce(func(d, o string) string {
			return d + o
		}, "hello ").(string)
		if ret != "hello world" {
			t.Fatal("expect 15 but get: ", ret)
		}
	})
}

func TestStreamCountComplex(t *testing.T) {
	t.Run("with string value", func(t *testing.T) {
		ret := newFunc("123", "456", "789").Filter(func(s string) bool {
			return s != "456"
		}).FlatMap(func(s string) []string {
			return strings.Split(s, "")
		}).Map(func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		}).Filter(func(i int) bool {
			if i == 2 || i == 7 {
				return false
			}
			return true
		}).Count()

		if ret != 4 {
			t.Fatal("expect 4 but get: ", ret)
		}
	})
}

func TestStreamForeachComplex(t *testing.T) {
	t.Run("with string value", func(t *testing.T) {
		newFunc("5646", "3221", "7789").Filter(func(s string) bool {
			return s != "5646"
		}).FlatMap(func(s string) []string {
			return strings.Split(s, "")
		}).Peek(func(s string) {
			t.Logf("after flatmap: %s\n", s)
		}).Map(func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		}).Peek(func(d int) {
			t.Logf("after map: %d\n", d)
		}).Sort(func(a, b int) int {
			return a - b
		}).Peek(func(v int) {
			t.Logf("after Sort: %d\n", v)
		}).Distinct(func(a, b int) bool {
			return a == b
		}).Peek(func(v int) {
			t.Logf("after distincet: %d\n", v)
		}).Filter(func(i int) bool {
			if i == 2 || i == 7 {
				return false
			}
			return true
		}).Peek(func(v int) {
			t.Logf("after fiter: %d\n", v)
		}).Foreach(func(i int) {
			t.Log("foreach ", i)
		})
	})
}
