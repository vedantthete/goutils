package collections

import (
	"fmt"
)

type Set[T comparable] struct {
	container map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{container: map[T]struct{}{}}
}

func (s *Set[T]) Length() int {
	return len(s.container)
}

func (s *Set[T]) Add(element T) {
	s.container[element] = struct{}{}
}

func (s *Set[T]) Remove(element T) error {
	_, exists := s.container[element]
	if exists {
		delete(s.container, element)
		return nil
	}
	return &NotFoundError{fmt.Sprintf("%v does not exist in the set", element)}
}

func (s *Set[T]) Discard(element T) {
	delete(s.container, element)
}

func (s *Set[T]) Clear() {
	for k := range s.container {
		delete(s.container, k)
	}
}

func (s *Set[T]) Copy() *Set[T] {
	setCopy := NewSet[T]()
	for k := range s.container {
		setCopy.container[k] = struct{}{}
	}
	return &setCopy
}

func (s *Set[T]) Difference(sets ...Set[T]) Set[T] {
	diffSet := NewSet[T]()
	for k := range s.container {
		exists := false
		for _, set := range sets {
			_, exists = set.container[k]
			if exists {
				break
			}
		}
		if !exists {
			diffSet.Add(k)
		}
	}
	return diffSet
}

func (s *Set[T]) Intersection(sets ...Set[T]) Set[T] {
	intersectionSet := NewSet[T]()
	for k := range s.container {
		exists := true
		for _, set := range sets {
			_, exists = set.container[k]
			if !exists {
				break
			}
		}
		if exists {
			intersectionSet.Add(k)
		}
	}
	return intersectionSet
}

func (s *Set[T]) IsDisjoint(set Set[T]) bool {
	for k := range s.container {
		_, exists := set.container[k]
		if exists {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsSubSet(set Set[T]) bool {
	for k := range s.container {
		_, exists := set.container[k]
		if !exists {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsSuperSet(set Set[T]) bool {
	for k := range set.container {
		_, exists := s.container[k]
		if !exists {
			return false
		}
	}
	return true
}

func (s *Set[T]) SymmetricDifference(set Set[T]) Set[T] {
	symmetricDiffSet := NewSet[T]()
	for k := range s.container {
		_, exists := set.container[k]
		if !exists {
			symmetricDiffSet.Add(k)
		}
	}
	for k := range set.container {
		_, exists := s.container[k]
		if !exists {
			symmetricDiffSet.Add(k)
		}
	}
	return symmetricDiffSet
}

func (s *Set[T]) Union(sets ...Set[T]) Set[T] {
	unionSet := NewSet[T]()
	for k := range s.container {
		unionSet.Add(k)
	}
	for _, set := range sets {
		for k := range set.container {
			unionSet.Add(k)
		}
	}
	return unionSet
}

func (s *Set[T]) Update(sets ...Set[T]) {
	for _, set := range sets {
		for k := range set.container {
			s.Add(k)
		}
	}
}
