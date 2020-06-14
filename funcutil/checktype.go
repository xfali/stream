// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package funcutil

import "reflect"

func VerifySlice(o interface{}) bool {
	if reflect.TypeOf(o).Kind() == reflect.Slice {
		return true
	}
	return false
}
