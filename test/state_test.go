// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/stream/valve"
	"testing"
)

func TestState(t *testing.T) {
	v := valve.SetState(valve.NORMAL, valve.SORTED)
	if !valve.CheckState(v, valve.SORTED) {
		t.Fatal("expect sorted but get false")
	}

	v = valve.SetState(v, valve.DISTINCT)
	if !valve.CheckState(v, valve.SORTED) || !valve.CheckState(v, valve.DISTINCT) {
		t.Fatal("expect sorted and distinct but get false")
	}

	v = valve.UnsetState(v, valve.SORTED)
	if valve.CheckState(v, valve.SORTED) {
		t.Fatal("expect not sorted but get false")
	}
	if !valve.CheckState(v, valve.DISTINCT) {
		t.Fatal("expect distinct but get false")
	}

	v = valve.UnsetState(v, valve.DISTINCT)
	if valve.CheckState(v, valve.SORTED) {
		t.Fatal("expect not sorted but get false")
	}
	if valve.CheckState(v, valve.DISTINCT) {
		t.Fatal("expect not distinct but get false")
	}
}
