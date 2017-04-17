package basic

import (
	"sort"
	"sync"
)

type SetInt struct {
	m map[int]bool
	sync.RWMutex
}

func NewSetInt() *SetInt {
	return &SetInt{
		m: map[int]bool{},
	}
}

func (s *SetInt) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *SetInt) Remove(item int) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

func (s *SetInt) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *SetInt) Len() int {
	return len(s.List())
}

func (s *SetInt) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
}

func (s *SetInt) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *SetInt) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := [](int){}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func (s *SetInt) SortList() []int {
	s.RLock()
	defer s.RUnlock()
	list := [](int){}
	for item := range s.m {
		list = append(list, item)
	}
	sort.Ints(list)
	return list
}
