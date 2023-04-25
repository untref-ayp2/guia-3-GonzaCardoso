package slicelist

//import (
//	"errors"
//	"fmt"
//)

// Implementar la lista sobre slices
type slice[T comparable] struct {
	slice []T
}

func NewSliceList[T comparable]() *slice[T] {
	return &slice[T]{}
}
func (s *slice[T]) Append(value T) {
	s.slice = append(s.slice, value)
}

func (s *slice[T]) Prepend(value T) {
	s.slice = append([]T{value}, s.slice...)
}

func (s *slice[T]) Get(i int) T {
	if i < 0 || i > len(s.slice) {
		var t T
		return t
	}
	return s.slice[i]
}
func (s *slice[T]) Size() int {
	return len(s.slice)
}
