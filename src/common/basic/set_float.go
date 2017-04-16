package basic

import (
	"sort"
	"sync"
)

type SetFLoat struct {
	m map[float64]bool
	sync.RWMutex
}

func NewSetFloat() *SetFLoat {
	return &SetFLoat{
		m: map[float64]bool{},
	}
}

func (s *SetFLoat) Add(item float64) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *SetFLoat) Remove(item float64) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

func (s *SetFLoat) Has(item float64) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *SetFLoat) Len() int {
	return len(s.List())
}

func (s *SetFLoat) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[float64]bool{}
}

func (s *SetFLoat) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *SetFLoat) List() []float64 {
	s.RLock()
	defer s.RUnlock()
	list := [](float64){}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func (s *SetFLoat) SortList() []float64 {
	s.RLock()
	defer s.RUnlock()
	list := [](float64){}
	for item := range s.m {
		list = append(list, item)
	}
	sort.Float64s(list)
	return list
}
