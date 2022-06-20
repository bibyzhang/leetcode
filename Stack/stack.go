//stack
package stack

import (
	"fmt"
	"sync"
)

type item interface {}

type stack struct {
	items []item
	lock sync.RWMutex
}

func NewStack() *stack {
	s := &stack{}
	s.items = []item{}
	return s
}

func (s *stack) Push (t item) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, t)
}

func (s *stack) Pop() item {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.items) == 0 {
		return nil
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}

func (s *stack) Print ()  {
	fmt.Println(s.items)
}

func (s *stack) Empty () bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}