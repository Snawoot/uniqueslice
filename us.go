package uniqueslice

import (
	"reflect"
	"unique"
)

// Make returns a globally unique handle for a slice of values of type T. Handles
// are equal if and only if the values used to produce them are equal.
// Make is safe for concurrent use by multiple goroutines.
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

// Handle is a globally unique identity for some slice of value of type T.
//
// Two handles compare equal exactly if the two values used to create the handles
// would have also compared equal. The comparison of two handles is trivial and
// typically much more efficient than comparing the values used to create them.
type Handle[Slice ~[]T, T comparable] struct {
	h unique.Handle[any]
}

// Value returns a slice of shallow copies of the T value that produced the Handle.
// Value is safe for concurrent use by multiple goroutines.
func (h Handle[Slice, T]) Value() Slice {
	vi := h.h.Value()
	rv := reflect.ValueOf(vi)
	rt := rv.Type()
	res := reflect.Indirect(reflect.New(rt))
	reflect.Copy(res, rv)
	return res.Slice(0, rt.Len()).Interface().(Slice)
}
