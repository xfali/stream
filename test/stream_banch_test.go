// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package test

import (
	"github.com/xfali/stream"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

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
			stream.Slice(benchSlice).Count()
		}
	})
}

func BenchmarkPipelineCount(b *testing.B) {
	benchSlice := makeSlice()
	b.Run("pipeline", func(b *testing.B) {
		for i:=0; i<b.N; i++ {
			stream.Slice(benchSlice).Filter(func(s string) bool {
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
			stream.Slice(benchSlice).Filter(func(s string) bool {
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
}

