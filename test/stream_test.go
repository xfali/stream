// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package test

import (
	"github.com/xfali/stream"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

var newFunc = stream.Pipeline

func TestStreamFilter(t *testing.T) {
	s := newFunc([]int{1, 2, 3, 4, 5})
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
	s := newFunc([]int{1, 2, 3, 4, 5})
	o := s.FindFirst()
	if !reflect.DeepEqual(1, o.Get()) {
		t.Fatalf("Stream.First() failed: expected %v got %v", 1, o.Get())
	}
}

func TestStreamLast(t *testing.T) {
	s := newFunc([]int{1, 2, 3, 4, 5})
	o := s.FindLast()
	if !reflect.DeepEqual(5, o.Get()) {
		t.Fatalf("Stream.First() failed: expected %v got %v", 1, o.Get())
	}
}

func TestStreamFindAny(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := newFunc([]int{1, 2, 3, 4, 5})
		t.Log(s.FindAny().Get())
	}
	for i := 0; i < 10; i++ {
		s := newFunc([]int{1, 2, 3, 4, 5})
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
	s := newFunc([]int{1, 2, 3, 4, 5})
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
		newFunc([]int{1, 2, 3, 4, 5}).Skip(2).Foreach(func(i int) {
			t.Log(i)
			if i == 1 || i == 2 {
				t.Fatal("Skip 1、2 but got it")
			}
		})
	})

	t.Run("skip 0", func(t *testing.T) {
		newFunc([]int{1, 2, 3, 4, 5}).Skip(0).Foreach(func(i int) {
			switch i {
			case 1, 2, 3, 4, 5:
				t.Log(i)
			default:
				t.Fatal("cannot be here")
			}
		})
	})

	t.Run("skip 5", func(t *testing.T) {
		newFunc([]int{1, 2, 3, 4, 5}).Skip(5).Foreach(func(i int) {
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
	s := newFunc([]int{1, 2, 3, 4, 5})
	s.Skip(2).Limit(1).Foreach(func(i int) {
		t.Log(i)
		if i != 3 {
			t.Fatal("Skip 1、2 but got it")
		}
	})
}

func TestStreamDistinct(t *testing.T) {
	t.Run("distinct", func(t *testing.T) {
		newFunc([]int{1, 2, 2, 4, 5}).Distinct(func(a, b int) bool {
			return a == b
		}).Foreach(func(i int) {
			t.Log(i)
		})
	})

	t.Run("distinct twice", func(t *testing.T) {
		newFunc([]int{1, 2, 2, 4, 5}).Distinct(func(a, b int) bool {
			return a == b
		}).Distinct(func(a, b int) bool {
			return a == b
		}).Foreach(func(i int) {
			t.Log(i)
		})
	})

	t.Run("distinct after sort", func(t *testing.T) {
		newFunc([]int{1, 2, 2, 4, 5}).Sort(func(a, b int) int {
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
		s := newFunc([]int{5, 2, 3, 1, 4})
		s.Sort(func(a, b int) int {
			return a - b
		}).Foreach(func(i int) {
			t.Log(i)
		})
	})

	t.Run("desc", func(t *testing.T) {
		s := newFunc([]int{5, 2, 3, 1, 4})
		s.Sort(func(a, b int) int {
			return b - a
		}).Foreach(func(i int) {
			t.Log(i)
		})
	})

	t.Run("asc repeat", func(t *testing.T) {
		s := newFunc([]int{5, 2, 5, 2, 3, 1, 4})
		s.Sort(func(a, b int) int {
			return a - b
		}).Foreach(func(i int) {
			t.Log(i)
		})
	})

	t.Run("desc twice", func(t *testing.T) {
		s := newFunc([]int{5, 2, 3, 1, 4})
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
	newFunc([]string{"1", "2", "3"}).Map(func(s string) int {
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

func TestStreamFlatMap(t *testing.T) {
	newFunc([]string{"hello world", "xfali stream"}).FlatMap(func(s string) []string {
		return strings.Split(s, " ")
	}).Foreach(func(i string) {
		switch i {
		case "hello", "world", "xfali", "stream":
			t.Log(i)
		default:
			t.Fatal("cannot be here")
		}
	})

	newFunc([]string{"1,2,3,4", "5,6,7,8"}).FlatMap(func(s string) []int {
		return newFunc(strings.Split(s, ",")).Map(func(s string) int {
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
		c := newFunc([]int{1, 2, 3, 4, 5}).Limit(2).Count()
		if c != 2 {
			t.Fatal("expect 2 but get: ", c)
		}
	})

	t.Run("filter", func(t *testing.T) {
		c := newFunc([]int{1, 2, 3, 4, 5}).Filter(func(a int) bool {
			return a != 2
		}).Count()
		if c != 4 {
			t.Fatal("expect 2 but get: ", c)
		}
	})
}

func TestStreamForeach(t *testing.T) {
	t.Run("limit", func(t *testing.T) {
		newFunc([]int{1, 2, 3, 4, 5}).Limit(2).Foreach(func(i int) {
			switch i {
			case 1,2:
				t.Log(i)
			default:
				t.Fatal("cannot be here")
			}
		})
	})

	t.Run("filter", func(t *testing.T) {
		newFunc([]int{1, 2, 3, 4, 5}).Filter(func(a int) bool {
			return a != 2
		}).Foreach(func(i int) {
			switch i {
			case 1,3,4,5:
				t.Log(i)
			default:
				t.Fatal("cannot be here")
			}
		})
	})
}

func TestStreamPeek(t *testing.T) {
	t.Run("limit", func(t *testing.T) {
		c := newFunc([]int{1, 2, 3, 4, 5}).Limit(2).Peek(func(i int) {
			switch i {
			case 1,2:
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
		c := newFunc([]int{1, 2, 3, 4, 5}).Filter(func(a int) bool {
			return a != 2
		}).Peek(func(i int) {
			switch i {
			case 1,3,4,5:
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
		c := newFunc([]int{1, 2, 3, 4, 5}).AnyMatch(func(i int) bool {
			return i == 3
		})
		if !c {
			t.Fatal("expect true but get: ", c)
		}
	})

	t.Run("false", func(t *testing.T) {
		c := newFunc([]int{1, 2, 3, 4, 5}).AnyMatch(func(i int) bool {
			return i == 6
		})
		if c {
			t.Fatal("expect false but get: ", c)
		}
	})
}

func TestStreamAllMatch(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		c := newFunc([]int{1, 2, 3, 4, 5}).AllMatch(func(i int) bool {
			return i != 6
		})
		if !c {
			t.Fatal("expect true but get: ", c)
		}
	})

	t.Run("false", func(t *testing.T) {
		c := newFunc([]int{1, 2, 3, 4, 5}).AllMatch(func(i int) bool {
			return i != 4
		})
		if c {
			t.Fatal("expect false but get: ", c)
		}
	})
}

func TestStreamReduce(t *testing.T) {
	t.Run("without init value", func(t *testing.T) {
		ret := newFunc([]int{1, 2, 3, 4, 5}).Reduce(func(d, o int) int {
			return d + o
		}, nil).(int)
		if ret != 15 {
			t.Fatal("expect 15 but get: ", ret)
		}
	})

	t.Run("with init value", func(t *testing.T) {
		ret := newFunc([]int{1, 2, 3, 4, 5}).Reduce(func(d, o int) int {
			return d + o
		}, 2).(int)
		if ret != 17 {
			t.Fatal("expect 17 but get: ", ret)
		}
	})

	t.Run("with string init value", func(t *testing.T) {
		ret := newFunc([]string{"w", "o", "r", "l", "d"}).Reduce(func(d, o string) string {
			return d + o
		}, "hello ").(string)
		if ret != "hello world" {
			t.Fatal("expect 15 but get: ", ret)
		}
	})
}

func TestStreamCountComplex(t *testing.T) {
	t.Run("with string value", func(t *testing.T) {
		ret := newFunc([]string{"123", "456", "789"}).Filter(func(s string) bool {
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
		newFunc([]string{"5646", "3221", "7789"}).Filter(func(s string) bool {
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

func makeSlice() []string {
	var ret []string
	for i:=0; i<256; i++ {
		ret = append(ret, strconv.Itoa(rand.Intn(99999999)))
	}
	return ret
}

func BenchmarkPipelineSimpleCount(b *testing.B) {
	benchSlice := makeSlice()
	b.Run("pipeline", func(b *testing.B) {
		for i:=0; i<b.N; i++ {
			stream.Pipeline(benchSlice).Count()
		}
	})

	b.Run("simple_slice", func(b *testing.B) {
		for i:=0; i<b.N; i++ {
			stream.SimpleSlice(benchSlice).Count()
		}
	})
}

func BenchmarkPipelineCount(b *testing.B) {
	benchSlice := makeSlice()
	b.Run("pipeline", func(b *testing.B) {
		for i:=0; i<b.N; i++ {
			stream.Pipeline(benchSlice).Filter(func(s string) bool {
				return s != "5646"
			}).FlatMap(func(s string) []string {
				return strings.Split(s, "")
			}).Map(func(s string) int {
				i, _ := strconv.Atoi(s)
				return i
			}).Sort(func(a, b int) int {
				return a - b
			}).Distinct(func(a, b int) bool {
				return a == b
			}).Filter(func(i int) bool {
				if i == 2 || i == 7 {
					return false
				}
				return true
			}).Count()
		}
	})

	b.Run("simple_slice", func(b *testing.B) {
		for i:=0; i<b.N; i++ {
			stream.SimpleSlice(benchSlice).Filter(func(s string) bool {
				return s != "5646"
			}).FlatMap(func(s string) []string {
				return strings.Split(s, "")
			}).Map(func(s string) int {
				i, _ := strconv.Atoi(s)
				return i
			}).Sort(func(a, b int) int {
				return a - b
			}).Distinct(func(a, b int) bool {
				return a == b
			}).Filter(func(i int) bool {
				if i == 2 || i == 7 {
					return false
				}
				return true
			}).Count()
		}
	})
}

func BenchmarkPipelineForeach(b *testing.B) {
	benchSlice := makeSlice()
	b.Run("pipeline", func(b *testing.B) {
		for i:=0; i<b.N; i++ {
			stream.Pipeline(benchSlice).Filter(func(s string) bool {
				return s != "5646"
			}).FlatMap(func(s string) []string {
				return strings.Split(s, "")
			}).Map(func(s string) int {
				i, _ := strconv.Atoi(s)
				return i
			}).Sort(func(a, b int) int {
				return a - b
			}).Distinct(func(a, b int) bool {
				return a == b
			}).Filter(func(i int) bool {
				if i == 2 || i == 7 {
					return false
				}
				return true
			}).Foreach(func(i int) {
				//b.Log(i)
			})
		}
	})

	b.Run("simple_slice", func(b *testing.B) {
		for i:=0; i<b.N; i++ {
			stream.SimpleSlice(benchSlice).Filter(func(s string) bool {
				return s != "5646"
			}).FlatMap(func(s string) []string {
				return strings.Split(s, "")
			}).Map(func(s string) int {
				i, _ := strconv.Atoi(s)
				return i
			}).Sort(func(a, b int) int {
				return a - b
			}).Distinct(func(a, b int) bool {
				return a == b
			}).Filter(func(i int) bool {
				if i == 2 || i == 7 {
					return false
				}
				return true
			}).Foreach(func(i int) {

			})
		}
	})
}
