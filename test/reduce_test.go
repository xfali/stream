package test

import (
	"github.com/xfali/stream/funcutil"
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	in := []int{1, 2, 3}
	add := func(a, b int) int {
		return a + b
	}
	expect := 6
	out, err := funcutil.Reduce(add, in)
	if err != nil {
		t.Fatalf("Reduce() failed: %v", err)
	}
	if !reflect.DeepEqual(expect, out) {
		t.Fatalf("Reduce() failed: expect %v got %v", expect, out)
	}
}
