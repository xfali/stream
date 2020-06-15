// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package collection

import (
	"reflect"
)

type Collection interface {
	ElemType() reflect.Type
	Size() int
	Iterator() Iterator

	Add(value reflect.Value)
}

type Iterator interface {
	HasNext() bool
	Next() reflect.Value
}
