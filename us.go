package uniqueslice

import (
	"reflect"
	"unique"
)

func Make[Slice ~[]T, T comparable](s Slice) Handle[Slice, T] {
	l := len(s)
	arrTyp := reflect.ArrayOf(l, reflect.TypeFor[T]())
	arr := reflect.Indirect(reflect.New(arrTyp))
	view := arr.Slice(0, l).Interface().([]T)
	copy(view, s)
	return Handle[Slice, T]{
		h: unique.Make[any](arr.Interface()),
	}
}

type Handle[Slice ~[]T, T comparable] struct {
	h unique.Handle[any]
}

func (h Handle[Slice, T]) Value() Slice {
	vi := h.h.Value()
	rv := reflect.ValueOf(vi)
	rt := rv.Type()
	res := reflect.Indirect(reflect.New(rt))
	reflect.Copy(res, rv)
	return res.Slice(0, rt.Len()).Interface().(Slice)
}
