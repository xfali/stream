// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package stream

import (
    "errors"
    "reflect"
    sort2 "sort"
)

func Sort(function, slice interface{}) (ret interface{}, err error) {
    return sort(function, slice)
}

func sort(function, slice interface{}) (interface{}, error) {
    in := reflect.ValueOf(slice)
    if in.Kind() != reflect.Slice {
        return nil, errors.New("Limit The first param is not a slice ")
    }
    fn := reflect.ValueOf(function)
    inType := in.Type().Elem()
    if !verifyCompareFunction(fn, inType) {
        panic("sort: Function must be of type func(" + inType.String() + "," + inType.String() + ") int")
    }

    ss := sortSlice{
        v:       in,
        compare: fn,
        swap: reflect.Swapper(slice),
    }
    sort2.Sort(&ss)
    return ss.v.Interface(), nil
}

type sortSlice struct {
    v       reflect.Value
    compare reflect.Value
    swap    func(i, j int)
}

func (p *sortSlice) Len() int {
    return p.v.Len()
}

func (p *sortSlice) Less(i, j int) bool {
    var param [2]reflect.Value
    param[0] = p.v.Index(i)
    param[1] = p.v.Index(j)
    return p.compare.Call(param[:])[0].Int() < 0
}

func (p *sortSlice) Swap(i, j int) {
    p.swap(i, j)
}
