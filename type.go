package sets

import (
	"fmt"
)

type Set[T comparable] map[T]bool

func CreateSet[T comparable](els ...T) Set[T] {
	res := make(Set[T])

	for _, v := range els {
		res[v] = true
	}
	return res
}

func (set Set[T]) Add(el T) {
	set[el] = true
}

func (set Set[T]) Remove(el T) (err bool) {
	_, ok := set[el]
	if ok {
		delete(set, el)
	}
	return ok
}

func (set Set[T]) String() (res string) {
	res += "{"
	for el := range set {
		res += fmt.Sprint(el)
		res += ", "
	}
	res = res[:len(res)-2]
	res += "}"
	return
}

func (set Set[T]) Contains(el T) bool {
	_, ok := set[el]
	return ok
}

func (set Set[T]) Copy() (res Set[T]) {
	res = make(Set[T])
	for el := range set {
		res.Add(el)
	}
	return
}
